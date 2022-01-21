package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainType "JHE_admin/internal/types"
	"JHE_admin/model"
	"JHE_admin/web/hall/service"
	"JHE_admin/web/hall/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type HallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserLogic) GetNotice() (*mainType.Result, error) {
	var notice []mainType.Notice

	err := global.GVA_DB.Model(&notice).Order("create_time desc").Find(&notice).Error

	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: notice,
	}, nil
}
func (u *UserLogic) GetFeedBack(req types.FeedBack) (*mainType.Result, error) {
	var feed []types.FeedBack

	err := global.GVA_DB.Model(&feed).Where("uid = ? and u_is_read = ?", req.Uid, 0).Order("create_time desc").Find(&feed).Error

	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: feed,
	}, nil
}
func (u *UserLogic) ReadFeedBack(req types.FeedBack) (*mainType.Result, error) {
	err := global.GVA_DB.Model(&types.FeedBack{}).Where("id = ?", req.Id).Update("u_is_read", 1).Error

	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "阅读失败",
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "阅读成功",
	}, nil
}
func (u *UserLogic) GetRankList() (*mainType.Result, error) {

	list1, list2, list3, err := service.GetRankList()

	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	day, week, month, _, err := model.GetGameConfig()
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: &types.RankList{
			RankList: types.Rank{
				DayRank:   list1,
				WeekRank:  list2,
				MonthRank: list3,
			},
			Config: types.GameConfig{
				Day:   day,
				Week:  week,
				Month: month,
			},
		},
	}, nil
}
