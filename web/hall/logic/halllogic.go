package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainType "JHE_admin/internal/types"
	"JHE_admin/table"
	"JHE_admin/web/hall/service"
	"JHE_admin/web/hall/types"
	"context"
	"database/sql"
	"fmt"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
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

	err := global.GVA_DB.Model(&feed).Where("uid = ? and uis_read = ? and status = 1 ", req.Uid, 0).Order("create_time desc").Find(&feed).Error

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
	err := global.GVA_DB.Model(&types.FeedBack{}).Where("id = ?", req.Id).Update("uis_read", 1).Error

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
	day, week, month, err := service.GetGameConfig()
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
				Day:    day,
				Week:   week,
				Month:  month,
				Symbol: "JHE",
			},
		},
	}, nil
}
func (u *UserLogic) GetRechargeList(req types.CustomerPage) (*mainType.Result, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var total int64
	var list []types.UserRecharge
	err := global.GVA_DB.Table("user_recharges").Select("uid,date,symbol,amount,pre_balance,now_balance").Where("uid = ?", req.Uid).Count(&total).Limit(limit).Offset(offset).Order("date desc").Find(&list).Error
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败" + err.Error(),
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: types.PageResult{
			Total:    total,
			List:     list,
			Page:     req.Page,
			PageSize: req.Page,
		},
	}, nil
}
func (u *UserLogic) GetRewardList(req types.CustomerPage) (*mainType.Result, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var total int64
	var list []types.UserWithdrawl
	err := global.GVA_DB.Model(&list).Select("uid,finish_date,pre_balance,now_balance,amount").Where("uid = ? AND status = 2", req.Uid).Count(&total).Limit(limit).Offset(offset).Order("finish_date desc").Find(&list).Error
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败" + err.Error(),
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: types.PageResult{
			Total:    total,
			List:     list,
			Page:     req.Page,
			PageSize: req.Page,
		},
	}, nil
}
func (u *UserLogic) GetProfitList(req types.CustomerPage) (*mainType.Result, error) {
	uid := u.ctx.Value("uid").(int)
	total, list, err := service.GetProfitList(req, uid)
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败" + err.Error(),
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: types.PageResult{
			Total:    total,
			List:     list,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}, nil
}

func (u *UserLogic) Config() (*mainType.Result, error) {
	config, err := service.GetConfig()
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败" + err.Error(),
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: config,
	}, nil
}
func (u *UserLogic) VipPopulariseProfit() (*mainType.Result, error) {
	uid := u.ctx.Value("uid").(int)
	db := global.GVA_DB.WithContext(context.Background())
	var gameSumProfit sql.NullFloat64
	var sum float64
	var profitList []types.CustomerOperator
	err := db.Model(&profitList).Where("uid = ? and is_draw = 0", uid).Select("SUM(num) as num,symbol").Group("symbol").Find(&profitList).Error
	if err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	fmt.Println(sum)
	if gameSumProfit.Valid {
		sum = gameSumProfit.Float64
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: profitList,
	}, nil
}
func (u *UserLogic) GetCurrency() (*mainType.Result, error) {

	var currency []table.Currency
	err := global.GVA_DB.Model(&currency).Find(&currency).Error
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取失敗",
			Data: currency,
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: currency,
	}, nil
}
func (u *UserLogic) GetBanner() (*mainType.Result, error) {
	db := global.GVA_DB
	var fileLists []mainType.ExaFileUploadAndDownload
	err := db.Find(&fileLists).Order("updated_at desc").Find(&fileLists).Error
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "获取banner失败",
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "获取banner成功",
		Data: fileLists,
	}, nil
}
func (u *UserLogic) DrawProfit() (*mainType.Result, error) {

	uid := u.ctx.Value("uid").(int)
	err := service.DrawProfit(uid)
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "领取成功",
	}, nil
}
func (u *UserLogic) VipConfig() (*mainType.Result, error) {
	param := global.GVA_CacheSysConfig.GetSysParameter()
	var config = &types.VipConfig{}
	config.BuyVipCost = param.BuyVipCost
	config.BuyVipSymbol = param.BuyVipSymbol

	return &mainType.Result{
		Code: 200,
		Msg:  "获取参数成功",
		Data: config,
	}, nil
}
func (u *UserLogic) BuyVip() (*mainType.Result, error) {
	uid := u.ctx.Value("uid").(int)
	err := service.UserBuyVip(uid)
	if err != nil {
		return &mainType.Result{
			Code: 400,
			Msg:  "充值失败," + err.Error(),
		}, nil
	}
	return &mainType.Result{
		Code: 200,
		Msg:  "充值成功",
	}, nil
}
