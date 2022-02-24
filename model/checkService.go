/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 12:31:45
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-19 14:55:26
 */
package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"time"

	"github.com/araddon/dateparse"
)

func GetRewardList(req types.RewardReq) (total int64, list []types.UserWithdrawl, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GVA_DB.Model(&types.UserWithdrawl{})
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
			return 0, list, sum, err
		}
		if req.StartTime == req.EndTime {
			ad, _ := time.ParseDuration("24h")
			end = end.Add(ad)
		}
		db = db.Where("req_date BETWEEN ? AND ?", start, end)
	}

	if err = db.Count(&total).Error; err != nil {
		return total, list, sum, err
	}
	if err = db.Limit(limit).Offset(offset).Order("req_date desc").Find(&list).Error; err != nil {
		return total, list, sum, err
	}
	for k, _ := range list {
		sum += list[k].Amount
	}
	return total, list, sum, err
}

func GetDailyReward() (num int64, total float64, err error) {
	var start_date = time.Now().Format("2006-01-02")
	start, err := dateparse.ParseAny(start_date)
	ad, _ := time.ParseDuration("24h")
	end := start.Add(ad)

	db := global.GVA_DB.Model(&types.UserWithdrawl{}).Where("finish_date BETWEEN ? AND ?", start, end).Where("status = ?", "2")
	var sum []float64
	if err = db.Count(&num).Pluck("amount", &sum).Error; err != nil {
		return num, total, err
	}

	for k, _ := range sum {
		total = total + sum[k]
	}
	return num, total, err
}
