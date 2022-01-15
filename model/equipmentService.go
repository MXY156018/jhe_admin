package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"time"
)

func GetEquipment(req types.EquipmentList) (list []types.Equipments, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var equiList []types.Equipments
	db := global.GVA_DB.Model(&types.Equipments{}).Joins("left join driving_schools on equipments.driving_school = driving_schools.id")

	if req.DrivingSchool != 0 {
		db = db.Where("driving_schools.id = ?", req.DrivingSchool)
	}
	if req.Name != "" {
		db = db.Where("equipments.name LIKE ?", "%"+req.Name+"%")
	}
	if req.EId != "" {
		db = db.Where("equipments.e_id LIKE ?", "%"+req.EId+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return equiList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&equiList).Error
	return equiList, total, err
}
func CreateEquipment(req types.Equipments) error {

	req.CreateTime = time.Now()
	req.IsUpdate = 0
	if err := global.GVA_DB.Model(&req).Create(&req).Error; err != nil {
		return err
	}
	return nil
}
