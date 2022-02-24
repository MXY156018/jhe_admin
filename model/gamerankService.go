package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"context"
	"database/sql"
	"time"

	"github.com/araddon/dateparse"
	"go.uber.org/zap"
)

func GetGameConfig() (dayProfit float64, weekprofit float64, month float64, sumprofit float64, err error) {
	var gameConfig []types.GameRankConfig
	err = global.GVA_DB.Model(&gameConfig).Find(&gameConfig).Error
	if err != nil {
		return dayProfit, weekprofit, month, sumprofit, err
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

	sumprofit, err = GetSumProfit()
	return dayProfit, weekprofit, month, sumprofit, err
}

func GetRankList(req types.GameRankSearch) (list1 []*types.GameRankList, list2 []*types.GameRankList, list3 []*types.GameRankList, err error) {
	// var gameranklist []types.GameRankList
	var day_start, day_end, week_start, week_end, month_start, month_end string
	// var Str1, Str2 string
	if req.Time != "" {
		// var start_date = time.Now().Format("2006-01-02")
		// day_start, _ = dateparse.ParseAny(req.Time)
		// ad, _ := time.ParseDuration("24h")
		// day_end = day_start.Add(ad)

		// Str1, Str2 = utils.GetWeekDay(day_start, false)
		// week_start, _ = dateparse.ParseAny(Str1)
		// week_end, _ = dateparse.ParseAny(Str2)

		// Str1, Str2 := utils.GetMonthDay(day_start, false)
		// month_start, _ = dateparse.ParseAny(Str1)
		// month_end, _ = dateparse.ParseAny(Str2)
		start_date, _ := dateparse.ParseAny(req.Time)
		day_start, day_end = utils.GetToday(start_date, true)
		week_start, week_end = utils.GetWeekDay(start_date, true)
		month_start, month_end = utils.GetMonthDay(start_date, true)
	} else {
		// var start_date = time.Now().Format("2006-01-02")
		// day_start, _ = dateparse.ParseAny(start_date)
		// ad, _ := time.ParseDuration("24h")
		// day_end = day_start.Add(ad)

		// Str1, Str2 = utils.GetWeekDay(day_start, false)
		// week_start, _ = dateparse.ParseAny(Str1)
		// week_end, _ = dateparse.ParseAny(Str2)

		// Str1, Str2 := utils.GetMonthDay(day_start, false)
		// month_start, _ = dateparse.ParseAny(Str1)
		// month_end, _ = dateparse.ParseAny(Str2)
		day_start, day_end = utils.GetToday(time.Now(), true)
		week_start, week_end = utils.GetWeekDay(time.Now(), true)
		month_start, month_end = utils.GetMonthDay(time.Now(), true)
	}
	day_list := GetUserList(day_start, day_end, 1)
	week_list := GetUserList(week_start, week_end, 2)
	month_list := GetUserList(month_start, month_end, 3)
	if req.Uid != 0 {
		var usr []*types.GameRankList
		for k, v := range day_list {
			if v.Uid == req.Uid {
				v.Rank = k + 1
				usr = append(usr, v)
			}
		}
		day_list = usr
		for k, v := range week_list {
			if v.Uid == req.Uid {
				v.Rank = k + 1
				usr = append(usr, v)
			}
		}
		week_list = usr
		for k, v := range month_list {
			if v.Uid == req.Uid {
				v.Rank = k + 1
				usr = append(usr, v)
			}
		}
		month_list = usr
	} else {
		if len(day_list) > 0 {
			for k, v := range day_list {
				v.Rank = k + 1
			}
		}
		if len(week_list) > 0 {
			for k, v := range week_list {
				v.Rank = k + 1
			}
		}
		if len(month_list) > 0 {
			for k, v := range month_list {
				v.Rank = k + 1
			}
		}
	}
	return day_list, week_list, month_list, err
}

// func GetUserList(start_time time.Time, end_time time.Time, Type int64) []*types.GameRankList {
// 	var list []*types.GameRankList
// 	global.GVA_DB.Table("game_records").Select("SUM(score) as sum_score,uid,end_time").Where("end_time BETWEEN ? AND ?", start_time, end_time).Order("sum_score desc").Find(&list)

// 	db := global.GVA_DB.Table("customer_operators").Where("create_time BETWEEN ? AND ?", start_time, end_time)
// 	for _, v := range list {
// 		var sumlist []float64
// 		var sum float64
// 		v.StartTime = start_time
// 		if Type == 1 {
// 			db = db.Where("type = 4").Pluck("num", &sumlist)
// 			for _, v1 := range sumlist {
// 				sum = sum + v1
// 			}
// 			v.Profit = sum
// 		} else if Type == 2 {
// 			db = db.Where("type = 6").Pluck("num", &sumlist)
// 			for _, v1 := range sumlist {
// 				sum = sum + v1
// 			}
// 			v.EndTime = end_time
// 			v.Profit = sum
// 		} else {
// 			db = db.Where("type = 7").Pluck("num", &sumlist)
// 			for _, v1 := range sumlist {
// 				sum = sum + v1
// 			}
// 			v.EndTime = end_time
// 			v.Profit = sum
// 		}
// 	}
// 	return list
// }
func GetUserList(start_time string, end_time string, Type int64) []*types.GameRankList {
	var list []*types.GameRankList
	db := global.GVA_DB.WithContext(context.Background())
	if Type == 1 {
		db.Table("game_rank_today").Where("date = ?", start_time).Order("credit desc").Find(&list)
	} else if Type == 2 {
		db.Table("game_rank_week").Where("date = ?", start_time).Order("credit desc").Find(&list)
	} else if Type == 3 {
		db.Table("game_rank_month").Where("date = ?", start_time).Order("credit desc").Find(&list)
	}
	for _, v := range list {
		v.Start_time = start_time
		v.End_time = end_time

	}
	for k, _ := range list {
		var num sql.NullFloat64
		err := db.Table("customer_operators").Select("num").Where("uid = ? and create_time = ? and type <> 4", list[k].Uid, list[k].Start_time).Find(&num).Error
		if err == nil {
			if num.Valid {
				list[k].Profit = num.Float64
			}
		} else {
			global.GVA_LOG.Error("err", zap.Any("err", err))
		}

	}

	return list
}

func GetSumProfit() (sum float64, err error) {
	var num sql.NullFloat64
	err = global.GVA_DB.Table("customer_operators").Where("type = 1 or type = 2 or type =3").Select("SUM(num) as sumProfit").Find(&num).Error
	if err != nil {
		return 0, err
	}
	if num.Valid {
		sum = num.Float64
	}
	return sum, err
}
