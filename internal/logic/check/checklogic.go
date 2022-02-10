package check

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
		global.GVA_LOG.Error("獲取失敗!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: types.PageResult{
				Total:    total,
				List:     list,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
			Msg: "獲取失敗",
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
		Msg: "獲取成功",
	}, nil
}
func (c *CheckLogic) PassCheck(req types.Reward) (*types.Result, error) {
	req.CreateTime = time.Now()
	req.Status = 1
	err := global.GVA_DB.Model(&req).Create(&req).Error
	if err != nil {
		global.GVA_LOG.Error("提交失敗", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "提交失敗",
		}, nil
	}
	//申请接口转币

	return &types.Result{
		Code: 0,
		Msg:  "審核通過",
	}, nil
}
func (c *CheckLogic) PassCheckByIds(req types.Ids) (*types.Result, error) {

	return &types.Result{
		Code: 0,
		Msg:  "批量審核通過",
	}, nil
}
func (c *CheckLogic) GetDailyReward() (*types.Result, error) {
	num, total, err := model.GetDailyReward()
	if err != nil {
		global.GVA_LOG.Error("獲取失敗", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: types.DailyReward{
			Num:   num,
			Total: total,
		},
		Msg: "獲取成功",
	}, nil
}
