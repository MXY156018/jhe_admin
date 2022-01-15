package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
)

func GetDriveSchoolById(req types.DrivingSchool) (types.DrivingSchool, error) {
	var rep types.DrivingSchool
	if err := global.GVA_DB.Model(&rep).Where("id = ?", req.Id).Find(&req).Error; err != nil {
		return rep, err
	}
	return rep, nil
}
func GetDriveSchool(req types.DrivingSchool) ([]types.DrivingSchool, error) {
	var rep []types.DrivingSchool
	if err := global.GVA_DB.Model(&rep).Find(&rep).Error; err != nil {
		return rep, err
	}
	return rep, nil
}
