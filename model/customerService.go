package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func GetCustomerList(req types.CustimerSearch) (list []types.CustomerList, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var userList []types.CustomerList
	db := global.GVA_DB.Model(&types.Customers{}).Select("customers.id,customers.address,customers.type,customers.create_time,customers.status,wallet.balance").Joins("left join wallet on customers.id=wallet.uid")
	if req.Id != 0 {
		db = db.Where("customers.id LIKE ?", "%"+fmt.Sprint(req.Id)+"%")
	}
	if req.Status != "" {
		db = db.Where("customers.status = ?", req.Status)
	}
	if req.Type != "" {
		db = db.Where("customers.type = ?", req.Type)
	}
	if req.StartTime != "" && req.EndTime != "" {
		start, err := dateparse.ParseAny(req.StartTime)
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return list, total, err
		}
		if req.StartTime == req.EndTime {
			ad, _ := time.ParseDuration("24h")
			end = end.Add(ad)
		}
		db = db.Where("create_time BETWEEN ? AND ?", start, end)
	}
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Scan(&userList).Error

	return userList, total, err
}

func GetCustomerById(req types.Customers) (types.CustomerList, error) {
	var user types.CustomerList
	if req.Id <= 0 {
		return user, errors.New("参数错误")
	}
	db := global.GVA_DB.Model(&types.Customers{}).Select("customers.id,customers.address,customers.type,customers.create_time,customers.status,wallet.balance").Joins("left join wallet on customers.id=wallet.uid")

	if err := db.Where("id = ?", req.Id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetSubordinateModel(req types.CustomerList) (list []types.CustomerList, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var userList []types.CustomerList
	db := global.GVA_DB.Model(&types.Customers{}).Select("customers.id,customers.address,customers.type,customers.create_time,customers.status,wallet.balance").Joins("left join wallet on customers.id=wallet.uid")
	if req.Id != 0 {
		db = db.Where("customers.sid = ?", req.Id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Scan(&userList).Error

	return userList, total, err
}
func GetCustomerGameRecordModel(req types.GameRecordList) (list []types.GameRecord, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	if req.Uid <= 0 {
		return list, total, errors.New("参数错误")
	}
	db := global.GVA_DB.Model(&list).Where("uid = ?", req.Uid).Count(&total)
	if db.Error != nil {
		return list, total, db.Error
	}
	db = db.Limit(limit).Offset(offset).Scan(&list)
	if db.Error != nil {
		return list, total, db.Error
	}
	return list, total, nil
}
func GetCustomerOperatorModel(req types.OperateRecord) (list []types.CustomerOperator, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	if req.Uid <= 0 {
		return list, total, errors.New("参数错误")
	}
	db := global.GVA_DB.Model(&list).Where("uid = ?", req.Uid).Count(&total)
	if db.Error != nil {
		return list, total, db.Error
	}
	db = db.Limit(limit).Offset(offset).Scan(&list)
	if db.Error != nil {
		return list, total, db.Error
	}
	return list, total, nil
}
