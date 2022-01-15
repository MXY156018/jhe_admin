package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"time"

	"github.com/araddon/dateparse"
)

func GetRewardList(req types.RewardReq) (total int64, list []types.Reward, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GVA_DB.Model(&types.Reward{})
	if req.Uid != 0 {
		db = db.Where("uid = ?", req.Uid)
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.StartTime != "" && req.EndTime != "" {
		start, err := dateparse.ParseAny(req.StartTime)
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return 0, list, err
		}
		if req.StartTime == req.EndTime {
			ad, _ := time.ParseDuration("24h")
			end = end.Add(ad)
		}
		db = db.Where("end_time BETWEEN ? AND ?", start, end)
	}

	if err = db.Count(&total).Error; err != nil {
		return total, list, err
	}
	if err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error; err != nil {
		return total, list, err
	}
	return total, list, err
}

func GetDailyReward() (num int64, total float64, err error) {
	var start_date = time.Now().Format("2006-01-02")
	start, err := dateparse.ParseAny(start_date)
	ad, _ := time.ParseDuration("24h")
	end := start.Add(ad)

	db := global.GVA_DB.Model(&types.Reward{}).Where("create_time BETWEEN ? AND ?", start, end).Where("status = ?", "2")
	var sum []float64
	if err = db.Count(&num).Pluck("reward", &sum).Error; err != nil {
		return num, total, err
	}
	for _, v := range sum {
		total = total + v
	}
	return num, total, err
}
