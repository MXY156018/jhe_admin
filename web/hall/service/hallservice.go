package service

import (
	"JHE_admin/global"
	"JHE_admin/utils"
	"JHE_admin/web/hall/types"
	"time"

	"github.com/araddon/dateparse"
)

func GetRankList() (list1 []types.GameRankList, list2 []types.GameRankList, list3 []types.GameRankList, err error) {
	var day_start, day_end, week_start, week_end, month_start, month_end time.Time
	var Str1, Str2 string
	var start_date = time.Now().Format("2006-01-02")
	day_start, _ = dateparse.ParseAny(start_date)
	ad, _ := time.ParseDuration("24h")
	day_end = day_start.Add(ad)

	Str1, Str2 = utils.GetWeekDay(day_start)
	week_start, _ = dateparse.ParseAny(Str1)
	week_end, _ = dateparse.ParseAny(Str2)

	Str1, Str2 = utils.GetMonthDay(day_start)
	month_start, _ = dateparse.ParseAny(Str1)
	month_end, _ = dateparse.ParseAny(Str2)

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
func GetUserList(start_time time.Time, end_time time.Time, Type int64) (list []types.GameRankList) {
	global.GVA_DB.Table("game_records").Select("SUM(score) as sum_score,uid").Where("end_time BETWEEN ? AND ?", start_time, end_time).Group("uid").Order("sum_score desc").Find(&list)
	return list
}
