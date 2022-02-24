/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-22 15:48:22
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 18:22:48
 */
package systemconfig

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/table"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type SystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemConfigLogic {
	return SystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (s *SystemConfigLogic) GetSystemConfig() (*types.Result, error) {
	db := global.GVA_DB.WithContext(context.Background())
	params := []table.SysConfig{}
	// var config = make(map[string]string)
	err := db.Model(&params).Select("param,value,remark").Find(&params).Error
	if err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "获取系統參數失敗",
		}, nil
	}
	// for _, v := range params {
	// 	config[v.Param] = v.Value
	// }

	return &types.Result{
		Code: 0,
		Data: params,
		Msg:  "获取成功",
	}, nil
}
func (s *SystemConfigLogic) SetSystemConfig(config []table.SysConfig) (*types.Result, error) {
	db := global.GVA_DB.WithContext(context.Background())
	err := global.GVA_CacheSysConfig.Update(db, config)
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "修改失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "修改成功",
	}, nil
}
