/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 13:55:40
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-20 16:01:15
 */
package check

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"JHE_admin/table"
	"JHE_admin/utils"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 钱包提币请求参数
type withdrawReq struct {
	AppUid   int    `json:"appUid"`
	Nonce    string `json:"nonce"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Coin     string `json:"coin"`
	OrderId  string `json:"orderId"`
	Sign     string `json:"sign"`
}

type withdrawResp struct {
	Code int `json:"code"`
}

const (
	// 新请求
	withdrawStatusNew int8 = iota
	//处理中
	withdrawStatusProcessing
	//成功
	_
	//失败
	withdrawStatusFail
	//失败退回
	withdrawStatusFallback
)

var withdrawIdLock sync.Mutex
var isLockId map[int]bool = map[int]bool{}

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (c *CheckLogic) GetCheckList(req types.RewardReq) (*types.Result, error) {
	total, list, sum, err := model.GetRewardList(req)

	if err != nil {
		global.GVA_LOG.Error("獲取失敗!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: types.PageResult{
				Total:    total,
				List:     list,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
			Msg: "獲取失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: types.PageResult{
			Total:    total,
			List:     list,
			Page:     req.Page,
			PageSize: req.PageSize,
			Sum:      sum,
		},
		Msg: "獲取成功",
	}, nil
}
func (c *CheckLogic) PassCheck(req types.UserWithdrawl) (*types.Result, error) {
	req.ReqDate = time.Now()
	req.Status = 1
	err := global.GVA_DB.Model(&req).Create(&req).Error
	if err != nil {
		global.GVA_LOG.Error("提交失敗", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "提交失敗",
		}, nil
	}
	//申请接口转币

	return &types.Result{
		Code: 0,
		Msg:  "審核通過",
	}, nil
}
func (c *CheckLogic) PassCheckByIds(req types.Ids) (*types.Result, error) {

	return &types.Result{
		Code: 0,
		Msg:  "批量審核通過",
	}, nil
}
func (c *CheckLogic) GetDailyReward() (*types.Result, error) {
	num, total, err := model.GetDailyReward()
	if err != nil {
		global.GVA_LOG.Error("獲取失敗", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: types.DailyReward{
			Num:   num,
			Total: total,
		},
		Msg: "獲取成功",
	}, nil
}

// 锁定订单
func lockWithdrawOrder(id int, isLock bool) bool {
	withdrawIdLock.Lock()
	defer withdrawIdLock.Unlock()
	if !isLock {
		delete(isLockId, id)
		return true
	}
	if isLockId[id] {
		return false
	}
	isLockId[id] = true
	return true
}

// 批准提币
func (c *CheckLogic) ApproveWithdraw(req *types.AssetApproveWithdrawReq) *types.Result {
	resp := &types.Result{}
	param := global.GVA_CacheSysConfig.GetSysParameter()
	if param == nil {
		resp.Code = 1
		resp.Msg = "服务器内部错误，参数错误"
		return resp
	}
	// 查找订单是否存在
	db := global.GVA_DB.WithContext(context.Background())
	order := &table.UserWithdrawl{}
	err := db.First(order, req.ID).Error
	if err != nil {
		resp.Code = 7
		resp.Msg = err.Error()
		return resp
	}
	if order.Status != withdrawStatusNew {
		resp.Code = 7
		resp.Msg = "订单状态错误"
		return resp
	}
	isOK := lockWithdrawOrder(req.ID, true)
	if !isOK {
		resp.Code = 7
		resp.Msg = "订单处理中"
		return resp
	}
	defer lockWithdrawOrder(req.ID, false)

	order.Status = withdrawStatusProcessing
	db = global.GVA_DB.WithContext(context.Background())
	err = db.Model(order).Update("status", withdrawStatusProcessing).Error
	if err != nil {
		resp.Code = 7
		resp.Msg = fmt.Sprintf("更新订单状态错误 %s", err.Error())
		c.Logger.Error(err)
		return resp
	}

	wReq := &withdrawReq{
		AppUid:   param.WalletAppId,
		Nonce:    fmt.Sprintf("%d", order.ID),
		To:       order.To,
		Quantity: fmt.Sprintf("%.4f", order.Amount),
		Coin:     order.Symbol,
		OrderId:  fmt.Sprintf("%d", order.ID),
		// 暂时不使用签名方式验证
		Sign: fmt.Sprintf("%d", order.ID),
	}
	wresp := &withdrawResp{}
	host := fmt.Sprintf("%s/api/bsc/withdrawl", param.WalletURL)
	err = utils.HttpJsonPost(host, wReq, resp)
	if err != nil || wresp.Code != 0 {
		db = global.GVA_DB.WithContext(context.Background())
		merr := db.Model(order).Update("status", withdrawStatusFail).Error
		if merr != nil {
			c.Logger.Error("更改提币状态错误", merr)
		}
		resp.Code = 7
		resp.Msg = "请求提币错误"
	}
	return resp
}

//退回
func (c *CheckLogic) BackWithdraw(req *types.AssetApproveWithdrawReq) *types.Result {
	resp := &types.Result{}
	// 查找订单是否存在
	db := global.GVA_DB.WithContext(context.Background())
	order := &table.UserWithdrawl{}
	err := db.First(order, req.ID).Error
	if err != nil {
		resp.Code = 7
		resp.Msg = err.Error()
		return resp
	}
	defer lockWithdrawOrder(req.ID, false)

	err = db.Transaction(func(tx *gorm.DB) error {
		// 更改状态
		merr := tx.Model(order).Update("status", 4).Error
		if merr != nil {
			return merr
		}
		// 更改余额
		merr = tx.Exec(fmt.Sprintf(
			"update wallets set withdraw_frozen=withdraw_frozen-%f,balance=balance+%f  where uid=%d and symbol='%s'",
			order.RawAmount, order.RawAmount, order.UID, order.Symbol,
		)).Error
		if merr != nil {
			return merr
		}
		return nil
	})
	if err != nil {
		resp.Code = 7
		resp.Msg = err.Error()
		return resp
	}
	return resp
}

// 提币失败退回
func (c *CheckLogic) WithdrawFallBack(req *types.AssetFallbackWithdrawReq) *types.Result {
	resp := &types.Result{}
	// 查找订单是否存在
	db := global.GVA_DB.WithContext(context.Background())
	order := &table.UserWithdrawl{}
	err := db.First(order, req.ID).Error
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp
	}
	if order.Status != withdrawStatusFail && order.Status != withdrawStatusNew {
		resp.Code = 2
		resp.Msg = "订单状态错误"
		return resp
	}
	isOK := lockWithdrawOrder(req.ID, true)
	if !isOK {
		resp.Code = 3
		resp.Msg = "订单处理中"
		return resp
	}
	defer lockWithdrawOrder(req.ID, false)

	err = db.Transaction(func(tx *gorm.DB) error {
		// 更改状态
		merr := tx.Model(order).Update("status", withdrawStatusFallback).Error
		if merr != nil {
			return merr
		}
		// 更改余额
		merr = tx.Exec(fmt.Sprintf(
			"update wallets set withdraw_frozen=withdraw_frozen-%f,balance=balance+%f  where uid=%d and symbol='%s'",
			order.RawAmount, order.RawAmount, order.UID, order.Symbol,
		)).Error
		if merr != nil {
			return merr
		}
		return nil
	})
	if err != nil {
		resp.Code = 4
		resp.Msg = err.Error()
		return resp
	}
	return resp
}

// 提币失败，已经手动处理，无需要返回冻结的额度
func (c *CheckLogic) WithdrawFailManulHandle(req *types.AssetWithdrawFailManulHandleReq) *types.Result {
	resp := &types.Result{}
	// 查找订单是否存在
	db := global.GVA_DB.WithContext(context.Background())
	order := &table.UserWithdrawl{}
	err := db.First(order, req.ID).Error
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp
	}
	if order.Status != withdrawStatusNew && order.Status != withdrawStatusFail {
		resp.Code = 2
		resp.Msg = "订单状态错误"
		return resp
	}
	isOK := lockWithdrawOrder(req.ID, true)
	if !isOK {
		resp.Code = 3
		resp.Msg = "订单处理中"
		return resp
	}
	defer lockWithdrawOrder(req.ID, false)

	err = db.Transaction(func(tx *gorm.DB) error {
		// 更改状态
		merr := tx.Model(order).Update("status", withdrawStatusFallback).Error
		if merr != nil {
			return merr
		}
		// 更改余额
		merr = tx.Exec(fmt.Sprintf(
			"update wallets set withdraw_frozen=withdraw_frozen-%f where uid=%d and symbol='%s'",
			order.RawAmount, order.UID, order.Symbol,
		)).Error
		if merr != nil {
			return merr
		}
		return nil
	})
	if err != nil {
		resp.Code = 4
		resp.Msg = err.Error()
		return resp
	}
	return resp
}
