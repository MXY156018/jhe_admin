/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-01-15 17:05:03
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 16:08:08
 */
package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainTypes "JHE_admin/internal/types"
	"JHE_admin/utils"
	"JHE_admin/web/hall/types"
	"context"
	"fmt"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserLogic) GetUserWallet() (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}
	var assets []types.Wallet
	db := global.GVA_DB.WithContext(context.Background())

	Uid := u.ctx.Value("uid")

	db = db.Table("wallets").Where("uid=?", Uid).Select("symbol,balance").Find(&assets)
	if db.Error != nil {
		result.Code = 400
		result.Msg = db.Error.Error()
		global.GVA_LOG.Error("err", zap.Any("err", db.Error))
		return result, nil
	}
	if len(assets) == 0 {
		result.Code = 400
		result.Msg = "用户资产不存在"
		return result, nil
	}
	result.Code = 200
	var address types.WalletAddress
	db.Table("user_block_chain_accounts").Where("uid=?", Uid).Limit(1).Select("address").Find(&address.Address)
	address.Wallet = assets
	result.Data = address
	return result, nil
}
func (u *UserLogic) GetInGameRank() (*mainTypes.Result, error) {
	var dayList, weekList, monthList []types.GameRankList
	var dayPre, weekPre, monthPre string
	var result = &mainTypes.Result{}
	dayPre = utils.GetPreDay(time.Now())
	weekPre = utils.GetPreWeek(time.Now())
	monthPre = utils.GetPreMonth(time.Now())

	fmt.Println(dayPre, weekPre, monthPre)
	db := global.GVA_DB.WithContext(context.Background())

	err := db.Table("game_rank_today as gt").Where("gt.date = ? and co.type=1 and co.create_time = ?", dayPre, dayPre).Joins("join customer_operators as co on gt.uid = co.uid").Select("gt.uid,gt.credit,co.num").Order("gt.credit desc").Find(&dayList).Error
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
		return result, nil
	}
	err = db.Table("game_rank_week as gw").Where("gw.date = ? and co.type = 2 and co.create_time = ?", weekPre, weekPre).Joins("join customer_operators as co on gw.uid = co.uid").Select("gw.uid,gw.credit,co.num").Order("gw.credit desc").Find(&weekList).Error
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
		return result, nil
	}
	err = db.Table("game_rank_month as gm").Where("gm.date = ? and co.type = 3 and co.create_time = ?", monthPre, monthPre).Joins("join customer_operators as co on gm.uid = co.uid").Select("gm.uid,gm.credit,co.num").Order("gm.credit desc").Find(&monthList).Error
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
		return result, nil
	}
	if len(dayList) > 0 {
		for k, _ := range dayList {
			dayList[k].Rank = k + 1

		}
	}
	if len(weekList) > 0 {
		for k, _ := range weekList {
			weekList[k].Rank = k + 1
		}
	}
	if len(monthList) > 0 {
		for k, _ := range monthList {
			monthList[k].Rank = k + 1
		}
	}
	return &mainTypes.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: types.Rank{
			DayRank:   dayList,
			MonthRank: monthList,
			WeekRank:  weekList,
		},
	}, nil
}
func (u *UserLogic) GetInGameRecord(req mainTypes.PageInfo) (*mainTypes.Result, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var total int64
	uid := u.ctx.Value("uid").(int)
	add, _ := time.ParseDuration("-144h")
	start_time := time.Now().Add(add).Format("2006-01-02") + " 00:00:00"
	end_time := time.Now().Format("2006-01-02") + " 23:59:59"

	fmt.Println(start_time, end_time)
	var gamerecord []mainTypes.GameRecordG1
	db := global.GVA_DB.WithContext(context.Background())
	db.Model(&gamerecord).Select("create_time,mode,win").Where("uid = ?", uid).Count(&total).Where("create_time >= ? and create_time < ?", start_time, end_time).Offset(offset).Limit(limit).Find(&gamerecord)
	if db.Error != nil {
		global.GVA_LOG.Error("err", zap.Any("err", db.Error))
		return &mainTypes.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	return &mainTypes.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: mainTypes.PageResult{
			List:     gamerecord,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}, nil
}
