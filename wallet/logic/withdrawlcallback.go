/*
 * @Descripttion:提币回调处理逻辑
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 00:28:30
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-19 18:13:13
 */
package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/wallet/types"
	"fmt"
	"strconv"
	"time"

	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type withdrawlOrder struct {
	OrderId   string
	UID       int
	Status    int8
	RawAmount float32
}

type WithdrawlCallbackLogic struct {
	Logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawlCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) WithdrawlCallbackLogic {
	return WithdrawlCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 提币回调
func (h *WithdrawlCallbackLogic) Callback(req *types.WithdrawlCallbackReq) (*types.CallbackResp, error) {
	var result *types.CallbackResp = &types.CallbackResp{}
	go h.onWithdrawl(req)
	return result, nil
}

// 处理提币
func (h *WithdrawlCallbackLogic) onWithdrawl(req *types.WithdrawlCallbackReq) {

	// 检查订单是否存在
	id, err := strconv.ParseInt(req.OrderId, 10, 64)
	if err != nil {
		h.Logger.Error("提币回调错误，订单错误", req)
		return
	}
	orders := []withdrawlOrder{}
	db := global.GVA_DB.WithContext(context.Background())
	db = db.Table("user_withdrawls").Where("id=?", id).Limit(1).Find(&orders)
	if db.Error != nil {
		h.Logger.Error("提币回调错误", db.Error, req)
		return
	}
	if len(orders) == 0 {
		h.Logger.Error("提币回调，订单不存在", req)
		return
	}
	order := &orders[0]
	if order.Status != 1 {
		h.Logger.Error("提币回调错误，订单已经处理", req)
		return
	}
	var status int8
	if req.Success == 0 {
		// 提币失败
		status = 3
	} else {
		status = 2
		// 更改冻结额度
		db = db.Exec(fmt.Sprintf(
			"update wallets set withdraw_frozen=withdraw_frozen-%f where uid=%d and symbol='%s'",
			order.RawAmount, order.UID, req.Symbol,
		))
		if db.Error != nil {
			h.Logger.Error("提币回调错误,更改用户余额错误", db.Error, req)
		}
	}
	nowStr := time.Now().Format("2006-01-02  15:04:05")
	db = global.GVA_DB.WithContext(context.Background())
	var now float32
	err = db.Table("wallets").Where("uid = ? and symbol = ?", order.UID, req.Symbol).Select("balance").Find(&now).Error
	if err != nil {
		h.Logger.Error("提币回调错误", err, req)
		return
	}
	db = global.GVA_DB.WithContext(context.Background())
	db = db.Exec(fmt.Sprintf(
		"update user_withdrawls set status=%d,finish_date='%s',tx_hash='%s',now_balance='%f' where id=%d",
		status, nowStr, req.Hash, now, id,
	))
	if db.Error != nil {
		h.Logger.Error("更改提币状态错误", db.Error, req)
	}
}
