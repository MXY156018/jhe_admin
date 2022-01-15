package service

import "JHE_admin/global"

func CheckHallUser(uid int64,token string)bool{
	var count int64
	if err:=global.GVA_DB.Table("hall_tokens").Where("uid = ? AND token = ?",uid,token).Count(&count).Error;err!=nil{
		return false
	}
	if count==0{
		return false
	}
	return true
}