/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 11:23:10
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-23 11:45:35
 */
package model

import (
	"JHE_admin/global"
	maintype "JHE_admin/internal/types"
	"JHE_admin/utils"
	"JHE_admin/web/hall/types"
	"time"
)

func GetUserTree(parent int) []*types.RechargeSum {
	treeList := []*types.RechargeSum{}
	user := GetSumRechargeWithInfo(parent)
	for _, v := range user {
		child := GetUserTree(v.Uid)
		node := &types.RechargeSum{
			Uid:         v.Uid,
			Parent:      v.Parent,
			Type:        v.Type,
			Status:      v.Status,
			SumRecharge: GetSumRecharge(v.Uid, true),
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

func GetSumRechargeWithInfo(parent int) []types.RechargeSum {
	var rechargeUser []types.RechargeSum
	global.GVA_DB.Table("users as u").Select("SUM(ur.amount) as sum_recharge,u.uid,u.type,u.parent,u.status").Joins("left join user_recharges as ur on u.uid=ur.uid").Where("u.parent = ?", parent).Group("u.uid").Find(&rechargeUser)
	return rechargeUser
}
func GetSumRechargeWithInfo1(parent int) []types.RechargeSum {
	var rechargeUser []types.RechargeSum
	global.GVA_DB.Table("users").Select("uid,type,parent,status").Where("parent = ?", parent).Find(&rechargeUser)
	return rechargeUser

}

//獲取用戶總充值
func GetSumRecharge(uid int, pre bool) float64 {
	amount := 0.00
	preStr := ""
	if pre {
		preStr = utils.GetPreDate(time.Now())
	} else {
		preStr = utils.GetSumDate()
	}
	global.GVA_DB.Table("user_sum_recharges").Select("amount").Where("uid = ? and time = ?", uid, preStr).Find(&amount)
	return amount
}

func GetSubSumRecharge(user *maintype.UserDetail) float64 {
	var sum float64
	userTree := GetUserTree(user.Uid)
	GetSubordinateSumRecharge(user, userTree)
	return sum
}

//獲取用戶下級總充值
func GetSubordinateSumRecharge(user *maintype.UserDetail, child []*types.RechargeSum) {
	if len(child) > 0 {
		for _, v := range child {
			user.SubSumRecharge += v.SumRecharge
			GetSubordinateSumRecharge(user, v.Children)
		}
	}
}
