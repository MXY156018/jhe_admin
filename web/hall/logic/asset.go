/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 21:03:35
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 13:52:01
 */
package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"fmt"

	// "JHE_admin/utils"
	mtypes "JHE_admin/internal/types"
	"JHE_admin/table"
	"JHE_admin/web/hall/types"
	"context"
	"sync"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

/**
appUid	number	是	用户appUid，可在后台查看
nonce	string	是	随机字符串
to	string	是	转到哪个地址
quantity	string	是	额度
coin	string	是	币种
orderId	string	是	用户自定义订单ID，最大长度为32
gasPrice	int	否	设置gasPrice，单位为GWei,如果不设置，将使用服务器设置的gasPrice
gasLimit	int	否	设置 gasLimit ,如果不设置，将使用服务器设置的 gasLimit
sign	string	是	签名
*/
// 钱包提币请求参数
// type withdrawReq struct {
// 	AppUid   int    `json:"appUid"`
// 	Nonce    string `json:"nonce"`
// 	To       string `json:"to"`
// 	Quantity string `json:"quantity"`
// 	Coin     string `json:"coin"`
// 	OrderId  string `json:"orderId"`
// 	Sign     string `json:"sign"`
// }

// type withdrawResp struct {
// 	Code int `json:"code"`
// }

const (
	// 新请求
	withdrawStatusNew int8 = iota
	//处理中
	// withdrawStatusProcessing
	//成功
	// withdrawStatusSuccess
	// _
	//失败
	// withdrawStatusFail
	//失败退回
	// withdrawStatusFallback
)

var withdrawMapLock sync.Mutex

// 用户是否正在提币
var isUserWithdraw map[int]bool = map[int]bool{}

// var withdrawIdLock sync.Mutex
// var isLockId map[int]bool = map[int]bool{}

type AssetLogic struct {
	Logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) AssetLogic {
	return AssetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetLogic) lockUser(uid int, isLock bool) bool {
	withdrawMapLock.Lock()
	defer withdrawMapLock.Unlock()
	if !isLock {
		delete(isUserWithdraw, uid)
		return true
	}
	if isUserWithdraw[uid] {
		return false
	}
	isUserWithdraw[uid] = true
	return true
}

// 提币请求
// 只写入记录，需要审核
func (l *AssetLogic) Withdraw(req *types.AssetWithdrawReq) (*mtypes.Result, error) {
	result := &mtypes.Result{}
	param := global.GVA_CacheSysConfig.GetSysParameter()
	if param == nil {
		result.Code = 400
		result.Msg = "服务器内部错误，参数错误"
		return result, nil
	}
	currency := global.GVA_CacheCurrency.GetCurrency(req.Symbol)
	if currency == nil {
		result.Code = 400
		result.Msg = "币种不存在"
		return result, nil
	}
	if currency.IsExp != 0 {
		result.Code = 400
		result.Msg = "体验币不能提币"
		return result, nil
	}
	if req.Amount < currency.MinWithdraw {
		result.Code = 400
		result.Msg = fmt.Sprintf("最小提币额度为 %f", currency.MinWithdraw)
		return result, nil
	}
	if req.Amount > currency.MaxWithdraw {
		result.Code = 400
		result.Msg = fmt.Sprintf("最大提币额度为 %f", currency.MaxWithdraw)
		return result, nil
	}
	uid := l.ctx.Value("uid").(int)

	var pre float32
	err := global.GVA_DB.Table("wallets").Where("uid = ? and symbol = ?", uid, req.Symbol).Select("balance").Find(&pre).Error
	if err != nil {
		result.Code = 400
		result.Msg = "服务器内部错误，获取余额失败"
		return result, nil
	}
	if pre < req.Amount {
		result.Code = 400
		result.Msg = "余额不足"
		return result, nil
	}
	isOK := l.lockUser(uid, true)
	if !isOK {
		result.Code = 200
		result.Msg = "正在提币"
		return result, nil
	}
	defer l.lockUser(uid, false)

	fee := req.Amount * param.WithdrawFee

	record := &table.UserWithdrawl{
		UID:        uid,
		Status:     withdrawStatusNew,
		Symbol:     req.Symbol,
		RawAmount:  req.Amount,
		Fee:        fee,
		Amount:     req.Amount - fee,
		To:         req.To,
		PreBalance: pre,
	}
	db := global.GVA_DB.WithContext(context.Background())
	err = db.Transaction(func(tx *gorm.DB) error {
		// 插入新的记录

		merr := tx.Omit("req_date", "finish_date", "tx_hash").Create(record).Error
		if merr != nil {
			return merr
		}
		// 更改余额
		merr = tx.Exec(fmt.Sprintf(
			"update wallets set withdraw_frozen=withdraw_frozen+%f,balance=balance-%f  where uid=%d and symbol='%s'",
			record.RawAmount, record.RawAmount, record.UID, record.Symbol,
		)).Error
		if merr != nil {
			return merr
		}
		return nil
	})

	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
		l.Logger.Error("提币请求错误", err.Error(), req)
		return result, nil
	}
	result.Code = 200
	result.Msg = "提币申请成功，等待管理员审核"
	return result, nil
}
