package gamerank

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type GameRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) GameRankLogic {
	return GameRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (g *GameRankLogic) GetGameRankConfig() (*types.Result, error) {
	day, week, month, sumprofit, err := model.GetGameConfig()
	if err != nil {
		global.GVA_LOG.Error("獲取失敗", zap.Any("err", err))
	}

	return &types.Result{
		Code: 0,
		Msg:  "獲取成功",
		Data: &types.GameConfig{
			Day:       day,
			Week:      week,
			Month:     month,
			SumProfit: sumprofit,
		},
	}, nil
}
func (g *GameRankLogic) SetGameRankConfig(req types.GameConfig) (*types.Result, error) {
	if req.Day <= 0 || req.Week <= 0 || req.Month <= 0 {
		return &types.Result{
			Code: 7,
			Msg:  "參數錯誤",
		}, nil
	}
	err := global.GVA_DB.Table("game_rank_configs").Where("id = 31").Update("num", req.Day).Error
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "設置失敗",
		}, nil
	}
	err = global.GVA_DB.Table("game_rank_configs").Where("id = 32").Update("num", req.Week).Error
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "設置失敗",
		}, nil
	}
	err = global.GVA_DB.Table("game_rank_configs").Where("id = 33").Update("num", req.Month).Error
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "設置失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "設置成功",
	}, nil
}
func (g *GameRankLogic) GetGameRank(req types.GameRankSearch) (*types.Result, error) {
	list1, list2, list3, err := model.GetRankList(req)
	if err != nil {
		return &types.Result{
			Code: 7,
			Data: types.RankList{
				DayList:   list1,
				WeekList:  list2,
				MonthList: list3,
			},
			Msg: "获取失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Data: types.RankList{
			DayList:   list1,
			WeekList:  list2,
			MonthList: list3,
		},
		Msg: "获取成功",
	}, nil
}
