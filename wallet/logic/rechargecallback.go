/*
 * @Descripttion:充值回调处理逻辑
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 00:28:24
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-19 18:12:44
 */
package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/table"
	"JHE_admin/utils"
	"JHE_admin/wallet/types"
	"fmt"
	"strconv"
	"strings"
	"time"

	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type userAddress struct {
	UID     int `gorm:"primaryKey"`
	Address string
}

type UserSumRecharge struct {
	Uid    int     `json:"uid"`
	Time   string  `json:"time"`
	Amount float32 `json:"amount"`
}

type RechargeCallbackLogic struct {
	Logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRechargeCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) RechargeCallbackLogic {
	return RechargeCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 充值回调
func (h *RechargeCallbackLogic) Callback(req *types.RechargeCallbackReq) (*types.CallbackResp, error) {
	var result *types.CallbackResp = &types.CallbackResp{}
	go h.onRecharge(req)
	return result, nil
}

// 处理充值
func (h *RechargeCallbackLogic) onRecharge(req *types.RechargeCallbackReq) {
	// 如果币是从提币底子转过来的，有可能是归集，忽略
	param := global.GVA_CacheSysConfig.GetSysParameter()
	if param != nil {
		if strings.EqualFold(strings.ToLower(param.WalletWithdrawAddress), strings.ToLower(req.From)) {
			h.Logger.Info("原地址为提币地址", req.From)
			return
		}
	}
	db := global.GVA_DB.WithContext(context.Background())
	users := []userAddress{}
	db = db.Table("user_block_chain_accounts").Where("address=?", req.To).Limit(1).Find(&users)
	if db.Error != nil {
		h.Logger.Error("充值回调错误", db.Error, req)
		return
	}
	// 检查是否存在
	if len(users) == 0 {
		h.Logger.Error("充值回调错误,用户不存在", db.Error, req)
		return
	}
	var feeRate float32 = 0
	if param != nil {
		feeRate = param.RechargeFee
	}

	rawAmount, _ := strconv.ParseFloat(req.Quantity, 32)
	user := &users[0]
	fee := rawAmount * float64(feeRate)
	amount := rawAmount - fee
	var pre float32
	db = global.GVA_DB.WithContext(context.Background())
	err := db.Table("wallets").Where("uid = ? and symbol = ?", user.UID, req.Symbol).Select("balance").Find(&pre).Error
	if err != nil {
		h.Logger.Error("获取钱包余额错误", db.Error, req)
	}
	// 更改余额
	db = db.Exec(fmt.Sprintf(
		"insert into wallets(uid,symbol,balance) values(%d,'%s',%f) on duplicate key update  balance=balance+%f",
		user.UID, req.Symbol, amount, amount,
	))
	if db.Error != nil {
		h.Logger.Error("充值回调错误,更改用户余额错误", db.Error, req)

	}
	var now float32
	db = global.GVA_DB.WithContext(context.Background())
	err = db.Table("wallets").Where("uid = ? and symbol = ?", user.UID, req.Symbol).Select("balance").Find(&now).Error
	if err != nil {
		h.Logger.Error("获取钱包余额错误", db.Error, req)

	}
	// 记录
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	record := &table.UserRecharge{}
	record.UID = user.UID
	record.Date = nowStr
	record.Symbol = req.Symbol
	record.Amount = float32(amount)
	record.TxHash = req.Hash
	record.RawAmount = float32(rawAmount)
	record.Fee = float32(fee)
	record.PreBalance = pre
	record.NowBalance = now
	db = global.GVA_DB.WithContext(context.Background())
	db = db.Create(record)
	if db.Error != nil {
		h.Logger.Error("保存充值记录错误", db.Error, req)
	}
	go h.onSumRecharge(*record)
}

//插入充值记录  累加到当前日期节点的总充值
func (h *RechargeCallbackLogic) onSumRecharge(record table.UserRecharge) {
	nowStr := utils.GetSumDate()
	db := global.GVA_DB.WithContext(context.Background())
	err := db.Exec(fmt.Sprintf(
		"insert into user_sum_recharges(uid,time,amount) values(%d,'%s',%f) on duplicate key update  amount=amount+%f",
		record.UID, nowStr, record.Amount, record.Amount,
	)).Error
	if err != nil {
		h.Logger.Error("累加用户总充值失败" + err.Error())
		return
	}
}
