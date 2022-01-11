package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
)

func NoticeList(req types.NoticePage) (total int64, list []types.Notice, err error) {

	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GVA_DB.Model(&types.Notice{})
	err = db.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error
	return total, list, err
}
