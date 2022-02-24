/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 17:32:48
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-23 16:22:09
 */
package service

import (
	"JHE_admin/global"
	"JHE_admin/table"
	"JHE_admin/utils"
	"JHE_admin/web/hall/types"
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetRankList() (list1 []types.GameRankList, list2 []types.GameRankList, list3 []types.GameRankList, err error) {
	var day_start, day_end, week_start, week_end, month_start, month_end string
	day_start, day_end = utils.GetToday(time.Now(), true)
	week_start, week_end = utils.GetWeekDay(time.Now(), true)
	month_start, month_end = utils.GetMonthDay(time.Now(), true)

	day_list := GetUserList(day_start, day_end, 1)
	week_list := GetUserList(week_start, week_end, 2)
	month_list := GetUserList(month_start, month_end, 3)
	if len(day_list) > 0 {
		for k, _ := range day_list {
			day_list[k].Rank = k + 1
		}
	}
	if len(week_list) > 0 {
		for k, _ := range week_list {
			week_list[k].Rank = k + 1
		}
	}
	if len(month_list) > 0 {
		for k, _ := range month_list {
			month_list[k].Rank = k + 1
		}
	}
	return day_list, week_list, month_list, err
}
func GetUserList(start_time, end_time string, Type int64) (list []types.GameRankList) {
	db := global.GVA_DB.WithContext(context.Background())
	if Type == 1 {
		db.Table("game_rank_today").Where("date = ?", start_time).Order("credit desc").Find(&list)
	} else if Type == 2 {
		db.Table("game_rank_week").Where("date = ?", start_time).Order("credit desc").Find(&list)
	} else if Type == 3 {
		db.Table("game_rank_month").Where("date = ?", start_time).Order("credit desc").Find(&list)
	}
	return list
}

func GetProfitList(req types.CustomerPage, uid int) (total int64, list []types.CustomerOperator, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	err = global.GVA_DB.Table("customer_operators").Select("create_time,num,uid,type,sum_num").Where("uid = ?", uid).Count(&total).Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error
	if err != nil {
		return total, list, err
	}
	return total, list, err

}

func GetConfig() (types.SysConfig, error) {
	var recharge, withdraw, game_cost, game_fee string

	var config types.SysConfig

	db := global.GVA_DB.WithContext(context.Background())

	err := db.Model(&config).Where("param = ?", "recharge_fee").Select("value").Find(&recharge).Error
	if err != nil {
		return config, err
	}
	err = db.Model(&config).Where("param = ?", "withdraw_fee").Select("value").Find(&withdraw).Error
	if err != nil {
		return config, err
	}
	err = db.Model(&config).Where("param = ?", "game_cost").Select("value").Find(&game_cost).Error
	if err != nil {
		return config, err
	}
	err = db.Model(&config).Where("param = ?", "game_fee").Select("value").Find(&game_fee).Error
	if err != nil {
		return config, err
	}
	config.RechargeFee = recharge
	config.WithdrawlFee = withdraw
	config.GameCost = game_cost
	config.GameFee = game_fee
	return config, err
}
func GetPopulariseProfit(uid int) (float32, error) {
	// predatestr := utils.GetPreDate(time.Now())
	// userTree := model.GetUserTree(uid)

	return 0, nil
}

func GetGameConfig() (dayProfit float64, weekprofit float64, month float64, err error) {
	var gameConfig []types.GameRankConfig
	err = global.GVA_DB.Model(&gameConfig).Find(&gameConfig).Error
	if err != nil {
		return dayProfit, weekprofit, month, err
	}
	for i := 0; i < len(gameConfig); i++ {
		if gameConfig[i].Id == 31 {
			dayProfit = gameConfig[i].Num
		}
		if gameConfig[i].Id == 32 {
			weekprofit = gameConfig[i].Num
		}
		if gameConfig[i].Id == 33 {
			month = gameConfig[i].Num
		}
	}
	return dayProfit, weekprofit, month, err
}

func DrawProfit(uid int) error {
	db := global.GVA_DB.WithContext(context.Background())
	var list []types.CustomerOperator
	err := db.Model(&list).Where("uid = ? and is_draw = 0", uid).Find(&list).Error
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return errors.New("没有可领取的奖励")
	}
	for _, v := range list {
		err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(fmt.Sprintf(
				"insert into wallets(uid,symbol,balance) values(%d,'%s',%f) on duplicate key update  balance=balance+%f",
				v.Uid, v.Symbol, v.Num, v.Num,
			)).Error; err != nil {
				return err
			}

			if err := tx.Table("customer_operators").Where("uid = ? and create_time = ? and symbol = ?", v.Uid, v.CreateTime, v.Symbol).Update("is_draw", 1).Update("draw_time", time.Now()).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func UserBuyVip(uid int) error {
	param := global.GVA_CacheSysConfig.GetSysParameter()
	db := global.GVA_DB.WithContext(context.Background())

	var wallet []table.Wallet
	if err := db.Model(&wallet).Where("uid = ? and symbol = ?", uid, param.BuyVipSymbol).Limit(1).Find(&wallet); err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
	}
	if len(wallet) == 0 {
		return errors.New("余额不足")
	} else if wallet[0].Balance < param.BuyVipCost {
		return errors.New("余额不足")
	}
	var user []types.User
	if err := db.Model(&user).Where("uid = ?", uid).Limit(1).Find(&user); err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
	}
	if len(user) == 0 {
		return errors.New("查无此人")
	} else if user[0].Type == "1" {
		return errors.New("已是VIP")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf(
			"insert into wallets(uid,symbol,balance) values(%d,'%s',%f) on duplicate key update  balance=balance-%f",
			uid, param.BuyVipSymbol, param.BuyVipCost, param.BuyVipCost,
		)).Error; err != nil {
			return err
		}
		var list = &types.UserBuyVip{}
		list.Uid = uid
		list.Symbol = param.BuyVipSymbol
		list.Num = param.BuyVipCost
		list.CreateTime = time.Now()
		if err := tx.Model(&list).Create(&list).Error; err != nil {
			return err
		}
		if err := tx.Exec(fmt.Sprintf(
			"insert into users(uid,type) values(%d,'%s') on duplicate key update  type='%s'",
			uid, "1", "1",
		)).Error; err != nil {
			return err
		}
		return nil
	})
}
