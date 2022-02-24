/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 11:00:21
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 18:27:32
 */
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
	uid := f.ctx.Value("uid")
	req.Uid = uid.(int)
	if err := global.GVA_DB.Model(&req).Create(&req).Error; err != nil {
		return &types.Result{
			Code: 400,
			Msg:  "反馈失败，请稍后再试",
		}, nil
	}
	return &types.Result{
		Code: 200,
		Msg:  "反馈成功",
	}, nil
}
