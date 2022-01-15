package bill

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type BillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBillLogic(ctx context.Context, svcCtx *svc.ServiceContext) BillLogic {
	return BillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (b *BillLogic) GetBillList(req types.BillReq) (*types.Result, error) {
	total, list, sum, err := model.GetBill(req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: &types.PageResult{
				List:     list,
				Total:    total,
				Sum:      sum,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
			Msg: err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: &types.PageResult{
			List:     list,
			Total:    total,
			Sum:      sum,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Msg: "获取成功",
	}, nil
}
func (b *BillLogic) GetGameBillList(req types.BillReq) (*types.Result, error) {
	total, list, sum, err := model.GetGameBill(req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: &types.PageResult{
				List:     list,
				Total:    total,
				Sum:      sum,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
			Msg: err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: &types.PageResult{
			List:     list,
			Total:    total,
			Sum:      sum,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Msg: "获取成功",
	}, nil
}
func (b *BillLogic) GetDailyBill() (*types.Result, error) {
	recharge, reward, platform, err := model.GetDailyBill()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: &types.GameDailyRep{
				Recharge: recharge,
				Reward:   reward,
				Platform: platform,
			},
			Msg: err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: &types.GameDailyRep{
			Recharge: recharge,
			Reward:   reward,
			Platform: platform,
		},
		Msg: "获取成功",
	}, nil
}
func (b *BillLogic) GetSumProfit() (*types.Result, error) {
	sum, err := model.GetSumPlatformProfit()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: sum,
			Msg:  err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: sum,
		Msg:  "获取成功",
	}, nil
}
