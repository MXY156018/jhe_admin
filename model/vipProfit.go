/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-21 11:35:48
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 10:32:17
 */
package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func CalculateVip() {
	var userId []int
	db := global.GVA_DB.WithContext(context.Background())

	err := db.Table("users").Where("type = 1").Select("uid").Find(&userId).Error
	if err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
	}
	var record = &types.CustomerOperator{}
	for _, v := range userId {
		userTree := GetUserTree(v)
		var user types.UserDetail
		GetSubordinateSumRecharge(&user, userTree)
		selfSum := GetSumRecharge(v, true)
		user.SumRecharge += selfSum
		sum := user.SubSumRecharge + user.SumRecharge
		profit := GetProfit(sum)
		var vipSum sql.NullFloat64
		var vipsum float64
		if profit > 0 {
			time := utils.GetPreDate(time.Now())
			if err := db.Table("customer_operators").Where("uid = ? and create_time < ? and symbol = ? and type = 1", v, time, "JHE").Select("SUM(num) as vipSum").Find(&vipSum); err != nil {
				global.GVA_LOG.Error("err", zap.Any("err", err))
			}
			if vipSum.Valid {
				vipsum = vipSum.Float64
			}
			err := db.Exec(fmt.Sprintf(
				"insert into customer_operators(uid,type,num,create_time,is_draw,symbol) values(%d,%d,%f,'%s',%d,'%s') on duplicate key update num=%f,sum_num=%f+%f",
				v, 4, profit, time, 0, "JHE", profit, profit, vipsum,
			)).Error
			if err != nil {
				global.GVA_LOG.Error("err"+fmt.Sprint(record.Uid)+fmt.Sprint(profit), zap.Any("err"+fmt.Sprint(record.Uid)+fmt.Sprint(profit), db.Error))
			}
		}
	}

}

func CalculateGameRank_Day(dayProfit float64) {
	preDay := utils.GetPreDay(time.Now())
	var daySum sql.NullFloat64
	var daysum float64
	day_list := GetUserList(preDay, "", 1)
	if len(day_list) > 0 {
		for k, v := range day_list {
			v.Rank = k + 1
		}
	}
	var config []types.GameRankConfig
	db := global.GVA_DB.WithContext(context.Background())
	db.Model(&config).Limit(30).Find(&config)
	if db.Error != nil {
		global.GVA_LOG.Error("err", zap.Any("err", db.Error))
	}
	for _, v := range day_list {
		for _, v1 := range config {
			if v.Rank == v1.Id {
				profit := dayProfit * v1.Num
				if err := db.Table("customer_operators").Where("uid = ? and create_time < ? and symbol = ? and type = 1", v.Uid, preDay, v1.Symbol).Select("SUM(num) as daySum").Find(&daySum); err != nil {
					global.GVA_LOG.Error("err", zap.Any("err", err))
				}
				if daySum.Valid {
					daysum = daySum.Float64
				}
				db.Exec(fmt.Sprintf(
					"insert into customer_operators(uid,type,num,create_time,is_draw,symbol,sum_num) values(%d,%d,%f,'%s',%d,'%s',%f) on duplicate key update num=%f,sum_num=%f",
					v.Uid, 1, profit, preDay, 0, v1.Symbol, daysum+profit, profit, daysum+profit,
				))
				break
			}
		}
	}
}

func CalculateGameRank_Week(weekprofit float64) {
	preWeek := utils.GetPreWeek(time.Now())
	var config []types.GameRankConfig
	db := global.GVA_DB.WithContext(context.Background())
	db.Model(&config).Limit(30).Find(&config)
	if db.Error != nil {
		global.GVA_LOG.Error("err", zap.Any("err", db.Error))
	}
	var weekSum sql.NullFloat64
	var weeksum float64
	week_list := GetUserList(preWeek, "", 2)
	if len(week_list) > 0 {
		for k, v := range week_list {
			v.Rank = k + 1
		}
	}
	for _, v := range week_list {
		for _, v1 := range config {
			if v.Rank == v1.Id {
				profit := weekprofit * v1.Num
				if err := db.Table("customer_operators").Where("uid = ? and create_time < ? and symbol = ? and type = 2", v.Uid, preWeek, v1.Symbol).Select("SUM(num) as weekSum").Find(&weekSum); err != nil {
					global.GVA_LOG.Error("err", zap.Any("err", err))
				}
				if weekSum.Valid {
					weeksum = weekSum.Float64
				}
				db.Exec(fmt.Sprintf(
					"insert into customer_operators(uid,type,num,create_time,is_draw,symbol,sum_num) values(%d,%d,%f,'%s',%d,'%s',%f) on duplicate key update  num=%f,sum_num=%f",
					v.Uid, 2, profit, preWeek, 0, v1.Symbol, weeksum+profit, profit, weeksum+profit,
				))
				break
			}
		}
	}
}

func CalculateGameRank_Month(monthprofit float64) {

	PreMonth := utils.GetPreMonth(time.Now())
	var config []types.GameRankConfig
	db := global.GVA_DB.WithContext(context.Background())
	db.Model(&config).Limit(30).Find(&config)
	if db.Error != nil {
		global.GVA_LOG.Error("err", zap.Any("err", db.Error))
	}
	var monthSum sql.NullFloat64
	var monthsum float64

	month_list := GetUserList(PreMonth, "", 3)

	if len(month_list) > 0 {
		for k, v := range month_list {
			v.Rank = k + 1
		}
	}

	for _, v := range month_list {
		for _, v1 := range config {
			if v.Rank == v1.Id {
				profit := monthprofit * v1.Num
				if err := db.Table("customer_operators").Where("uid = ? and create_time < ? and symbol = ? and type = 3", v.Uid, PreMonth, v1.Symbol).Select("SUM(num) as monthSum").Find(&monthSum); err != nil {
					global.GVA_LOG.Error("err", zap.Any("err", err))
				}
				if monthSum.Valid {
					monthsum = monthSum.Float64
				}
				db.Exec(fmt.Sprintf(
					"insert into customer_operators(uid,type,num,create_time,is_draw,symbol,sum_num) values(%d,%d,%f,'%s',%d,'%s',%f) on duplicate key update  num=%f,sum_num=%f",
					v.Uid, 3, profit, PreMonth, 0, v1.Symbol, monthsum+profit, profit, monthsum+profit,
				))
				break
			}
		}
	}
}

func GetGameConnfig() (dayProfit, weekprofit, month float64, err error) {
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

func GetProfit(sum float64) float64 {
	if sum >= 300000 {
		return sum * 0.04
	} else if sum >= 150000 {
		return sum * 0.035
	} else if sum >= 80000 {
		return sum * 0.03
	} else if sum >= 50000 {
		return sum * 0.025
	} else if sum >= 30000 {
		return sum * 0.02
	} else if sum >= 15000 {
		return sum * 0.015
	} else if sum >= 7000 {
		return sum * 0.01
	} else {
		return 0
	}
}
