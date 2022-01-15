package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainType "JHE_admin/internal/types"
	"JHE_admin/web/game/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type HallUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHallUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) HallUserLogic {
	return HallUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (h *HallUserLogic) UserLogin(req types.HallUser) (*mainType.Result, error) {
	if req.Uid <= 0 || req.Token == "" {
		return &mainType.Result{
			Code: 7,
			Msg:  "参数错误",
		}, nil
	}
	var count int64
	global.GVA_DB.Table("hall_tokens").Where("uid = ?", req.Uid).Count(&count)
	if count == 1 {
		if err := global.GVA_DB.Table("hall_tokens").Where("uid = ?", req.Uid).Update("token", req.Token).Error; err != nil {
			return &mainType.Result{
				Code: 7,
				Msg:  err.Error(),
			}, nil
		}
	} else {
		if err := global.GVA_DB.Table("hall_tokens").Create(req).Error; err != nil {
			return &mainType.Result{
				Code: 7,
				Msg:  err.Error(),
			}, nil
		}
	}
	return &mainType.Result{
		Code: 0,
		Msg:  "后台登录成功",
	}, nil
}
