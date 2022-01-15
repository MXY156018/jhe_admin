package model

import (
	"JHE_admin/global"
	maintype "JHE_admin/internal/types"
	"JHE_admin/web/hall/types"
)

func GetUserTree(sid int) []*types.RechargeSum {

	treeList := []*types.RechargeSum{}
	user := GetSumRecharge(sid)
	for _, v := range user {
		child := GetUserTree(v.Id)
		node := &types.RechargeSum{
			Id:          v.Id,
			Sid:         v.Sid,
			Type:        v.Type,
			Status:      v.Status,
			SumRecharge: v.SumRecharge,
			CreateTime:  v.CreateTime,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

func GetSumRecharge(sid int) (user []types.RechargeSum) {
	var rechargeUser []types.RechargeSum
	global.GVA_DB.Table("customers as c").Select("SUM(co.num) as sum_recharge,c.id,c.type,c.sid,c.status,co.create_time").Joins("left join customer_operators as co on c.id=co.uid").Where("c.sid = ? ", sid).Group("c.id").Find(&rechargeUser)
	return rechargeUser
}

func GetSubordinateSumRecharge(user *maintype.CustomerList, child []*types.RechargeSum) {
	if len(child) > 0 {
		for _, v := range child {
			user.SumSubordinateRecharge += v.SumRecharge
			GetSubordinateSumRecharge(user, v.Children)
		}
	}
}
