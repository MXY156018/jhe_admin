package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"context"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeedBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeedBackLogic {
	return FeedBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (f *FeedBackLogic) FeedBack(req types.FeedBack) (*types.Result, error) {
	req.CreateTime = time.Now()
	req.Status = "0"
	if err := global.GVA_DB.Model(&req).Create(&req).Error; err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "反馈失败，请稍后再试",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "反馈成功",
	}, nil
}
