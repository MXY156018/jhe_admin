package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"time"

	"github.com/araddon/dateparse"
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

	var sumprofitlist []float64
	err = global.GVA_DB.Table("customer_operators").Where("type = 4 or type = 6 or type = 7").Pluck("num", &sumprofitlist).Error
	for _, v := range sumprofitlist {
		sumprofit = sumprofit + v
	}
	return dayProfit, weekprofit, month, sumprofit, err
}

func GetRankList(req types.GameRankSearch) (list1 []*types.GameRankList, list2 []*types.GameRankList, list3 []*types.GameRankList, err error) {
	// var gameranklist []types.GameRankList
	var day_start, day_end, week_start, week_end, month_start, month_end time.Time
	var Str1, Str2 string
	if req.Time != "" {
		// var start_date = time.Now().Format("2006-01-02")
		day_start, _ = dateparse.ParseAny(req.Time)
		ad, _ := time.ParseDuration("24h")
		day_end = day_start.Add(ad)

		Str1, Str2 = utils.GetWeekDay(day_start)
		week_start, _ = dateparse.ParseAny(Str1)
		week_end, _ = dateparse.ParseAny(Str2)

		Str1, Str2 := utils.GetMonthDay(day_start)
		month_start, _ = dateparse.ParseAny(Str1)
		month_end, _ = dateparse.ParseAny(Str2)
	} else {
		var start_date = time.Now().Format("2006-01-02")
		day_start, _ = dateparse.ParseAny(start_date)
		ad, _ := time.ParseDuration("24h")
		day_end = day_start.Add(ad)

		Str1, Str2 = utils.GetWeekDay(day_start)
		week_start, _ = dateparse.ParseAny(Str1)
		week_end, _ = dateparse.ParseAny(Str2)

		Str1, Str2 := utils.GetMonthDay(day_start)
		month_start, _ = dateparse.ParseAny(Str1)
		month_end, _ = dateparse.ParseAny(Str2)
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

func GetUserList(start_time time.Time, end_time time.Time, Type int64) []*types.GameRankList {
	var list []*types.GameRankList
	global.GVA_DB.Table("game_records").Select("SUM(score) as sum_score,uid,end_time").Where("end_time BETWEEN ? AND ?", start_time, end_time).Order("sum_score desc").Find(&list)

	db := global.GVA_DB.Table("customer_operators").Where("create_time BETWEEN ? AND ?", start_time, end_time)
	for _, v := range list {
		var sumlist []float64
		var sum float64
		v.StartTime = start_time
		if Type == 1 {
			db = db.Where("type = 4").Pluck("num", &sumlist)
			for _, v1 := range sumlist {
				sum = sum + v1
			}
			v.Profit = sum
		} else if Type == 2 {
			db = db.Where("type = 6").Pluck("num", &sumlist)
			for _, v1 := range sumlist {
				sum = sum + v1
			}
			v.EndTime = end_time
			v.Profit = sum
		} else {
			db = db.Where("type = 7").Pluck("num", &sumlist)
			for _, v1 := range sumlist {
				sum = sum + v1
			}
			v.EndTime = end_time
			v.Profit = sum
		}
	}
	return list
}
