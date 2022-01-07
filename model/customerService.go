package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"errors"
	"fmt"
)

func GetCustomerList(req types.CustomerList) (list []types.CustomerList, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var userList []types.CustomerList
	db := global.GVA_DB.Model(&types.Customers{}).Select("customers.id,customers.address,customers.type,customers.create_time,customers.status,wallet.balance").Joins("left join wallet on customers.id=wallet.uid")
	if req.Id != 0 {
		db = db.Where("customers.id LIKE ?", "%"+fmt.Sprint(req.Id)+"%")
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
