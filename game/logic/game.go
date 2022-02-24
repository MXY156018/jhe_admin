/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 22:57:17
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 15:37:21
 */
// 游戏相关的API 接口
package logic

import (
	"JHE_admin/game/types"
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainTypes "JHE_admin/internal/types"
	"JHE_admin/table"
	"JHE_admin/utils"
	"fmt"
	"time"

	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

type GameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) GameLogic {
	return GameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 冻结额度
func (h *GameLogic) FreezeAssetForGame(req *types.GameFreezeAssetReq) (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}
	if req.IsFreeze {
		h.freezeAsset(req, result)
	} else {
		h.unfreezeAsset(req, result)
	}

	return result, nil
}

// 冻结
func (h *GameLogic) freezeAsset(req *types.GameFreezeAssetReq, result *mainTypes.Result) {
	uids := []int{}
	for i := 0; i < len(req.Items); i++ {
		uids = append(uids, req.Items[i].UID)
	}
	wallets := []table.Wallet{}

	db := global.GVA_DB.WithContext(context.Background())
	err := db.Where("uid in ?", uids).Where("symbol=?", req.Symbol).Find(&wallets).Error
	if err != nil {
		result.Code = 1
		result.Msg = err.Error()
		return
	}
	if len(wallets) != len(req.Items) {
		result.Code = 2
		result.Msg = "有用户额度不足"
		return
	}

	balances := []types.GameFreezeResp{}
	var target *table.Wallet
	for i := 0; i < len(req.Items); i++ {
		item := &req.Items[i]
		target = nil
		for j := 0; i < len(wallets); j++ {
			w := &wallets[j]
			if w.UID == item.UID {
				target = w
				break
			}
		}
		if target == nil {
			result.Code = 3
			result.Msg = "有用户额度不足"
			return
		}
		if target.Balance < item.Amount {
			result.Code = 4
			result.Msg = "有用户额度不足"
			return
		}
		// 余额
		balances = append(balances, types.GameFreezeResp{
			UID:    item.UID,
			Amount: target.Balance - item.Amount,
		})
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(req.Items); i++ {
			item := &req.Items[i]
			sql := fmt.Sprintf(
				"update wallets set game_frozen=game_frozen+%.2f,balance=balance-%.2f where uid=%d and symbol='%s'",
				item.Amount, item.Amount, item.UID, req.Symbol,
			)
			tx = tx.Exec(sql)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	if err != nil {
		result.Code = 5
		result.Msg = err.Error()
		h.Logger.Error("冻结余额错误", err, req)
	}

	result.Data = balances
}

// 解冻
func (h *GameLogic) unfreezeAsset(req *types.GameFreezeAssetReq, result *mainTypes.Result) {
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(req.Items); i++ {
			item := &req.Items[i]
			sql := fmt.Sprintf(
				"update wallets set game_frozen=game_frozen-%.2f,balance=balance+%.2f where uid=%d and symbol='%s'",
				item.Amount, item.Amount, item.UID, req.Symbol,
			)
			tx = tx.Exec(sql)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	if err != nil {
		result.Code = 1
		result.Msg = err.Error()
		h.Logger.Error("冻结余额错误", err, req)
	}
}

// 游戏结算
func (h *GameLogic) Settle(req *types.GameSettleReq) (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}
	// 通过事务方式保存
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(req.Items); i++ {
			// 更改用户余额
			item := &req.Items[i]
			sql := fmt.Sprintf(
				"update wallets set game_frozen=game_frozen-%.2f,balance=balance+%.2f where uid=%d and symbol='%s'",
				item.Unfreeze, item.Win, item.UID, req.Symbol,
			)
			tx = tx.Exec(sql)
			if tx.Error != nil {
				h.Logger.Error("游戏结算错误", tx.Error, item)
			}
		}
		return nil
	})

	go h.onSettle(req)
	if err != nil {
		result.Code = 1
		result.Msg = err.Error()
		h.Logger.Error("游戏结算错误", err, req)
		return result, nil
	}

	return result, nil
}

// 其他记录
func (h *GameLogic) onSettle(data *types.GameSettleReq) {
	db := global.GVA_DB.WithContext(context.Background())

	now := time.Now()
	today, _ := utils.GetToday(now, true)
	week, _ := utils.GetWeekDay(now, true)
	month, _ := utils.GetMonthDay(now, true)
	// 抽水 ...
	var commission float32 = 0
	var win float32
	var err error
	for i := 0; i < len(data.Items); i++ {
		item := &data.Items[i]
		// 抽水
		commission += item.Rebate
		// 投注 > 赢分 才有积分
		// Unfreeze 字段就是押分
		win = item.Win - item.Unfreeze
		if win > 0 {
			err = db.Exec(fmt.Sprintf(
				"insert into game_credit_today(`date`,`uid`,`credit`) values('%s',%d,%f) ON DUPLICATE KEY UPDATE credit=credit+%f",
				today, item.UID, win, win,
			)).Error
			if err != nil {
				h.Logger.Error("保存日积分记录错误", err, data)
			}

			err = db.Exec(fmt.Sprintf(
				"insert into game_credit_week(`date`,`uid`,`credit`) values('%s',%d,%f) ON DUPLICATE KEY UPDATE credit=credit+%f",
				week, item.UID, win, win,
			)).Error
			if err != nil {
				h.Logger.Error("保存周积分记录错误", err, data)
			}

			err = db.Exec(fmt.Sprintf(
				"insert into game_credit_month(`date`,`uid`,`credit`) values('%s',%d,%f) ON DUPLICATE KEY UPDATE credit=credit+%f",
				month, item.UID, win, win,
			)).Error
			if err != nil {
				h.Logger.Error("保存月积分记录错误", err, data)
			}

			err = db.Exec(fmt.Sprintf(
				"insert into game_rank_today(`date`,`uid`,`game_id`,`credit`) values('%s',%d,%d,%f) ON DUPLICATE KEY UPDATE credit=credit+%f",
				today, item.UID, data.GameId, win, win,
			)).Error
			if err != nil {
				h.Logger.Error("保存日排行记录错误", err, data)
			}

			err = db.Exec(fmt.Sprintf(
				"insert into game_rank_week(`date`,`uid`,`game_id`,`credit`) values('%s',%d,%d,%f) ON DUPLICATE KEY UPDATE credit=credit+%f",
				week, item.UID, data.GameId, win, win,
			)).Error
			if err != nil {
				h.Logger.Error("保存周排行记录错误", err, data)
			}

			err = db.Exec(fmt.Sprintf(
				"insert into game_rank_month(`date`,`uid`,`game_id`,`credit`) values('%s',%d,%d,%f) ON DUPLICATE KEY UPDATE credit=credit+%f",
				month, item.UID, data.GameId, win, win,
			)).Error
			if err != nil {
				h.Logger.Error("保存月排行记录错误", err, data)
			}
		}

		// 抽水
		if item.Rebate > 0 {
			err = db.Exec(fmt.Sprintf(
				"insert into game_commission(`uid`,`symbol`,`game_id`,`commission`) values('%d','%s',%d,%f) ON DUPLICATE KEY UPDATE commission=commission+%f",
				item.UID, data.Symbol, data.GameId, commission, commission,
			)).Error
			if err != nil {
				h.Logger.Error("保存用户抽水记录错误", err, item)
			}
		}
	}

	// 抽水记录
	if commission > 0 {
		err := db.Exec(fmt.Sprintf(
			"insert into game_commission(`uid`,`symbol`,`game_id`,`commission`) values(0,'%s',%d,%f) ON DUPLICATE KEY UPDATE commission=commission+%f",
			data.Symbol, data.GameId, commission, commission,
		)).Error
		if err != nil {
			h.Logger.Error("保存平台抽水记录错误", err, data)
		}
	}
}
