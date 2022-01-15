package check

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (c *CheckLogic) GetCheckList(req types.RewardReq) (*types.Result, error) {
	total, list, err := model.GetRewardList(req)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: types.PageResult{
				Total:    total,
				List:     list,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
			Msg: "获取失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: types.PageResult{
			Total:    total,
			List:     list,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Msg: "获取成功",
	}, nil
}
func (c *CheckLogic) PassCheck(req types.Reward) (*types.Result, error) {

	return &types.Result{
		Code: 0,
		Msg:  "审核通过",
	}, nil
}
func (c *CheckLogic) PassCheckByIds(req types.Ids) (*types.Result, error) {

	return &types.Result{
		Code: 0,
		Msg:  "批量审核通过",
	}, nil
}
func (c *CheckLogic) GetDailyReward() (*types.Result, error) {
	num, total, err := model.GetDailyReward()
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "获取失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: types.DailyReward{
			Num: num,
			Total: total,
		},
		Msg:  "获取成功",
	}, nil
}
