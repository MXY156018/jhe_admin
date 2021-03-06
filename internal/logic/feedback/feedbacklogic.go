package feedback

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"context"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
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

func (f *FeedBackLogic) GetFeedBackList(req types.FeedBackReq) (*types.Result, error) {

	total, list, err := model.GetFeedBack(req)
	if err != nil {
		global.GVA_LOG.Error("獲取失敗!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "獲取成功",
		Data: &types.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}, nil
}
func (f *FeedBackLogic) GetNewFeedBack() (*types.Result, error) {
	var total int64
	var feedback types.FeedBack
	err := global.GVA_DB.Model(&feedback).Where("is_read = ?", "0").Count(&total).Error
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "獲取成功",
		Data: total,
	}, nil
}
func (f *FeedBackLogic) ReadNewFeedBack(req types.FeedBack) (*types.Result, error) {
	var rep types.FeedBack
	err := global.GVA_DB.Model(&rep).Where("id = ?", req.Id).Find(&rep).Error
	if err != nil {
		global.GVA_LOG.Error("獲取失敗!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}

	if rep.Status == "0" {
		if err := global.GVA_DB.Model(&types.FeedBack{}).Where("id = ?", req.Id).Update("is_read", 1).Error; err != nil {
			return &types.Result{
				Code: 7,
				Msg:  "已讀失敗",
			}, nil
		}
	}
	return &types.Result{
		Code: 0,
		Msg:  "成功",
		Data: rep,
	}, nil

}
func (f *FeedBackLogic) DealFeedBack(req types.FeedBack) (*types.Result, error) {
	if err := global.GVA_DB.Model(&types.FeedBack{}).Where("id = ?", req.Id).Update("handle", req.Handle).Update("resolve_time", time.Now()).Update("status", "1").Error; err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "失败",
		}, nil
	}

	return &types.Result{
		Code: 0,
		Msg:  "處理成功",
	}, nil
}
func (f *FeedBackLogic) DeleteFeedBack(req types.FeedBack) (*types.Result, error) {

	if err := global.GVA_DB.Model(&types.FeedBack{}).Where("id = ?", req.Id).Update("status", "2").Error; err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "删除成功",
		}, nil
	}

	return &types.Result{
		Code: 0,
		Msg:  "删除失敗",
	}, nil
}
