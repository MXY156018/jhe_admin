package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"time"

	"github.com/araddon/dateparse"
)

func GetBill(req types.BillReq) (total int64, list []types.CustomerOperator, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := global.GVA_DB.Model(&list)
	if req.Uid != 0 {
		db = db.Where("uid = ?", req.Uid)
	}
	if req.Type != 0 {
		db = db.Where("type = ?", req.Type)
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

	if err := db.Count(&total).Error; err != nil {
		return total, list, err
	}
	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error
	return total, list, err
}
