/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 18:25:23
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 12:26:03
 */
package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"JHE_admin/utils"

	"JHE_admin/table"
	"JHE_admin/web/hall/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

// 验证参数
type verifyParam struct {
	// 地址
	Address string `json:"address"`
	// 签名
	Sign string `json:"sign"`
	// 签名信息
	RawMsg string `json:"rawMsg"`
}

type verifyResp struct {
	//错误码
	Code int `json:"code"`
}

/*
appUid	number	是	用户appUid，可在后台查看
nonce	string	是	随机字符串
address	string	否	如果为空，则服务器创建一个新的ETH账户
privateKey	string	否	私钥，用于自动转出
type	string	否	账户类型，如果不填，则为充值监听账户类型
sign	string	是	签名，见签名说明
*/
// 钱包创建地址参数
type walletCreateAccountReq struct {
	// app id
	AppUid int `json:"appUid"`
	// 随机
	Nonce string `json:"nonce"`
	//签名
	Sign string `json:"sign"`
}

// 钱包创建地址回复
type walletCreateAccountResp struct {
	// 错误码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 数据
	Data struct {
		Address string `json:"address"`
	} `json:"data"`
}

type EthSignLoginLogic struct {
	Logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEthSignLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) EthSignLoginLogic {
	return EthSignLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EthSignLoginLogic) Login(req *types.EthSignLoginReq, ip string) *types.LoginResp {
	result := &types.LoginResp{
		Code: 200,
	}
	param := global.GVA_CacheSysConfig.GetSysParameter()
	if param == nil {
		result.Code = 1
		result.Message = "系统参数错误"
		return result
	}
	if param.EthSignVerifyURL == "" {
		result.Code = 2
		result.Message = "不支持的登录方式"
		return result
	}

	now := time.Now().Unix()
	// 10 分钟有效时间
	if int64(req.Timestamp+600) < now {
		result.Code = 3
		result.Message = "超时"
		return result
	}
	signMsg := fmt.Sprintf("%s%d", req.Message, req.Timestamp)
	verifyReq := &verifyParam{}
	verifyReq.Address = req.Address
	verifyReq.RawMsg = signMsg
	verifyReq.Sign = req.Sign
	resp := &verifyResp{}

	// 签名验证，依赖于 nodejs 程序
	err := utils.HttpJsonPost(param.EthSignVerifyURL, verifyReq, resp)
	if err != nil {
		result.Code = 4
		result.Message = "服务器内部错误"
		l.Logger.Error(err)
		return result
	}
	if resp.Code != 0 {
		result.Code = 5
		result.Message = "签名验证错误"
		return result
	}
	l.onLogin(req, result, ip)

	return result
}

func (l *EthSignLoginLogic) onLogin(req *types.EthSignLoginReq, result *types.LoginResp, ip string) {
	address := strings.ToLower(req.Address)
	// 检查地址是否存在
	db := global.GVA_DB.WithContext(context.Background())
	users := []table.User{}
	err := db.Where("address=?", address).Limit(1).Find(&users).Error
	if err != nil {
		result.Code = 400
		result.Message = err.Error()
		return
	}
	if len(users) == 0 {
		l.onRegister(req, result, ip)
		return
	}
	user := &users[0]
	if user.Status == 0 {
		result.Code = 400
		result.Message = "用户已被冻结，请联系管理员"
		return
	}

	onLoginSuccess(l.svcCtx, result, user)
	go afterLogin(user.UID, ip)
}

func (l *EthSignLoginLogic) onRegister(req *types.EthSignLoginReq, result *types.LoginResp, ip string) {
	// 为用户生成钱包地址
	param := global.GVA_CacheSysConfig.GetSysParameter()
	if param == nil {
		result.Code = 8
		result.Message = "服务器内部错误，参数错误"
		return
	}
	host := fmt.Sprintf("%s/api/bsc/account/create", param.WalletURL)
	creq := &walletCreateAccountReq{}
	creq.AppUid = param.WalletAppId
	creq.Nonce = fmt.Sprintf("%d", rand.Int())
	creq.Sign = fmt.Sprintf("%d", rand.Int())
	cresp := &walletCreateAccountResp{}
	err := utils.HttpJsonPost(host, creq, cresp)
	if err != nil {
		result.Code = 9
		result.Message = err.Error()
		return
	}
	if cresp.Code != 0 {
		result.Code = 10
		result.Message = cresp.Message
		return
	}

	db := global.GVA_DB.WithContext(context.Background())
	if req.Parent > 0 {
		// 检查上级
		users := []table.User{}
		err := db.Where("uid=?", req.Parent).Limit(1).Find(&users).Error
		if err != nil {
			result.Code = 8
			result.Message = "上级不存在"
			return
		}
	}
	user := &table.User{}
	user.Address = strings.ToLower(req.Address)
	user.Head = int8(1 + rand.Int()%10)
	user.Parent = req.Parent
	err = db.Select("head", "parent", "address").Create(user).Error
	if err != nil {
		result.Code = 9
		result.Message = err.Error()
		return
	}

	// 生成充值地址
	account := &table.UserBlockchainAccount{}
	account.UID = user.UID
	account.Address = cresp.Data.Address
	err = db.Model(&account).Create(account).Error
	if err != nil {
		l.Logger.Error("创建钱包错误 ", err)
	}

	var wallet = []table.Wallet{
		{
			UID:    user.UID,
			Symbol: "JHE",
		}, {
			UID:    user.UID,
			Symbol: "bean",
		},
	}
	err = db.Model(&wallet).Create(&wallet).Error
	if err != nil {
		l.Logger.Error("创建用户钱包错误 ", err)
	}
	onLoginSuccess(l.svcCtx, result, user)
	go afterLogin(user.UID, ip)
}
