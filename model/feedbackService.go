package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"time"

	"github.com/araddon/dateparse"
)

func GetFeedBack(req types.FeedBackReq) (total int64, list []types.FeedBack, err error) {

	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := global.GVA_DB.Model(&types.FeedBack{})
	if req.Uid != 0 {
		db = db.Where("uid = ?", req.Uid)
	}
	if req.Message != "" {
		db = db.Where("message LIKE ?", "%"+req.Message+"%")
	}
	if req.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+req.Phone+"%")
	}
	if req.Email != "" {
		db = db.Where("email LIKE ?", "%"+req.Email+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.Status != "2" {
		db = db.Where("status <> ?", "2")
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
		db = db.Where("create_time BETWEEN ? AND ?", start, end)
	}
	err = db.Count(&total).Error
	if err != nil {
		return total, list, err
	}

	err = db.Limit(limit).Offset(offset).Order("status asc,create_time desc").Find(&list).Error
	return total, list, err
}
