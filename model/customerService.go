package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	subTypes "JHE_admin/web/hall/types"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func GetCustomerList(req types.CustimerSearch) (list []types.User, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var userList []types.User
	db := global.GVA_DB.Table("users").Select("uid,address,type,register_time,last_login_time,status").Where("is_bot = 0")
	if req.Uid != 0 {
		db = db.Where("uid LIKE ?", "%"+fmt.Sprint(req.Uid)+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
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
		db = db.Where("registerTime BETWEEN ? AND ?", start, end)
	}
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Scan(&userList).Error

	return userList, total, err
}

func GetCustomerById(req types.UserDetail) (types.UserDetail, error) {
	var user types.UserDetail
	if req.Uid <= 0 {
		return user, errors.New("参数错误")
	}
	db := global.GVA_DB.Table("users").Select("uid,parent,register_time,status,type,address").Where("is_bot = 0")

	if err := db.Where("uid = ?", req.Uid).Find(&user).Error; err != nil {
		return user, err
	}
	user.SumRecharge = GetSelfSumRecharge(req.Uid)
	GetAllSubSUmRecharge(&user)
	return user, nil
}

func GetSubordinateModel(req types.CustimerSearch) (list []types.User, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var userList []types.User
	db := global.GVA_DB.Model(&types.User{}).Select("uid,address,type,register_time,status")
	if req.Uid != 0 {
		db = db.Where("parent = ?", req.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Scan(&userList).Error

	return userList, total, err
}
func GetCustomerGameRecordModel(req types.GameRecordList) (list []types.GameRecordG1, total int64, err error) {
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
	db := global.GVA_DB.Model(&list).Where("uid = ? and is_draw = 1", req.Uid).Count(&total)
	if db.Error != nil {
		return list, total, db.Error
	}
	db = db.Limit(limit).Offset(offset).Scan(&list)
	if db.Error != nil {
		return list, total, db.Error
	}
	return list, total, nil
}

func GetSelfSumRecharge(uid int) float64 {
	var Sum sql.NullFloat64
	var sum float64
	err := global.GVA_DB.Table("user_sum_recharges").Select("SUM(amount) as Sum").Where("uid = ?", uid).Find(&Sum).Error
	if err != nil {
		return 0
	}
	if Sum.Valid {
		sum = Sum.Float64
	}
	return sum
}

func GetAllSubSUmRecharge(user *types.UserDetail) {
	userTree := GetUserTreeAllDate(user.Uid)
	GetSubordinateSumRecharge(user, userTree)
}

//计算首页数据

func GetSum(Type int) (sum float64, err error) {
	var result sql.NullFloat64
	db := global.GVA_DB.WithContext(context.Background())
	if Type == 1 {
		err = db.Table("user_recharges").Select("SUM(amount) as sum").Find(&result).Error
	} else if Type == 2 {
		err = db.Table("user_withdrawls").Where("status = 2").Select("SUM(amount) as sum").Find(&result).Error
	} else if Type == 3 {
		err = db.Table("customer_operators").Where("type = 4").Select("SUM(num) as sum").Find(&result).Error
	} else {
		err = db.Table("customer_operators").Where("type = 1 OR type = 2 OR type = 3").Select("SUM(num) as sum").Find(&result).Error
	}
	if err == nil {
		if result.Valid {
			sum = result.Float64
		}
	}
	return sum, err
}

func GetUserTreeAllDate(parent int) []*subTypes.RechargeSum {
	treeList := []*subTypes.RechargeSum{}
	user := GetSumRechargeWithInfo1(parent)
	for _, v := range user {
		child := GetUserTreeAllDate(v.Uid)
		node := &subTypes.RechargeSum{
			Uid:         v.Uid,
			Parent:      v.Parent,
			Type:        v.Type,
			Status:      v.Status,
			SumRecharge: GetAllSumRecharge(v.Uid),
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

func GetAllSumRecharge(uid int) float64 {
	fmt.Println("#########################################", uid)
	amount := 0.00
	global.GVA_DB.Table("user_sum_recharges").Select("amount").Where("uid = ?", uid).Find(&amount)
	return amount
}
