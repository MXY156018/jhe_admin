package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
)

func CreateSysOperationRecord(sysOperationRecord types.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
