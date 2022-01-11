package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"time"

	"github.com/araddon/dateparse"
)

func GetBill(req types.BillReq) (total int64, list []types.CustomerOperator, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list2 []types.CustomerOperator
	db1 := global.GVA_DB.Model(&list)
	db2 := global.GVA_DB.Model(&list2)

	if req.Uid != 0 {
		db1 = db1.Where("uid = ?", req.Uid)
		db2 = db2.Where("uid = ?", req.Uid)
	}
	if req.Type != 0 {
		db1 = db1.Where("type = ?", req.Type)
		db2 = db2.Where("type = ?", req.Type)
	}
	if req.StartTime != "" && req.EndTime != "" {
		start, err := dateparse.ParseAny(req.StartTime)
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return 0, list, 0, err
		}
		if req.StartTime == req.EndTime {
			ad, _ := time.ParseDuration("24h")
			end = end.Add(ad)
		}
		db1 = db1.Where("create_time BETWEEN ? AND ?", start, end)
		db2 = db2.Where("create_time BETWEEN ? AND ?", start, end)
	}

	if err := db2.Count(&total).Error; err != nil {
		return total, list, 0, err
	}
	var sum1 []float64

	err = db1.Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error
	var sum2 = 0.00
	db2.Pluck("num", &sum1).Scan(&list2)
	for i := 0; i < len(sum1); i++ {
		sum2 = sum2 + sum1[i]
	}
	sum = sum2
	return total, list, sum, err
}
func GetGameBill(req types.BillReq) (total int64, list []types.GameRecord, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list2 []types.GameRecord
	db1 := global.GVA_DB.Model(&list)
	db2 := global.GVA_DB.Model(&list2)

	if req.Uid != 0 {
		db1 = db1.Where("uid = ?", req.Uid)
		db2 = db2.Where("uid = ?", req.Uid)
	}
	if req.Type == 4 {
		db1 = db1.Where("win = ?", 0)
		db2 = db2.Where("win = ?", 0)
	}
	if req.StartTime != "" && req.EndTime != "" {
		start, err := dateparse.ParseAny(req.StartTime)
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return 0, list, 0, err
		}
		if req.StartTime == req.EndTime {
			ad, _ := time.ParseDuration("24h")
			end = end.Add(ad)
		}
		db1 = db1.Where("end_time BETWEEN ? AND ?", start, end)
		db2 = db2.Where("end_time BETWEEN ? AND ?", start, end)
	}

	if err := db2.Count(&total).Error; err != nil {
		return total, list, 0, err
	}
	var sum1 []float64

	err = db1.Limit(limit).Offset(offset).Order("end_time desc").Find(&list).Error
	var sum2 = 0.00
	db2.Pluck("commission", &sum1).Scan(&list2)
	for i := 0; i < len(sum1); i++ {
		sum2 = sum2 + sum1[i]
	}
	sum = sum2
	return total, list, sum, err
}
func GetDailyBill() (recharge float64, reward float64, platform float64, err error) {
	var start_date = time.Now().Format("2006-01-02")
	start, err := dateparse.ParseAny(start_date)
	ad, _ := time.ParseDuration("24h")
	end := start.Add(ad)
	var (
		rechargeList []float64
		rewardList   []float64
		platformList []float64
	)
	err = global.GVA_DB.Model(&types.CustomerOperator{}).Where("type = 1").Where("create_time BETWEEN ? AND ?", start, end).Pluck("num", &rechargeList).Error
	if err != nil {
		return recharge, reward, platform, err
	}
	err = global.GVA_DB.Model(&types.CustomerOperator{}).Where("type = 2").Where("create_time BETWEEN ? AND ?", start, end).Pluck("num", &rewardList).Error
	if err != nil {
		return recharge, reward, platform, err
	}
	for _, v := range rechargeList {
		recharge = recharge + v
	}
	for _, v := range rewardList {
		reward = reward + v
	}
	for _, v := range platformList {
		platform = platform + v
	}
	return recharge, reward, platform, err
}
