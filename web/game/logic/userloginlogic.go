package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainTypes "JHE_admin/internal/types"
	"JHE_admin/model"
	"JHE_admin/web/game/types"
	"context"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type HallUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHallUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) HallUserLogic {
	return HallUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (h *HallUserLogic) UserRegister(req types.Customer) (*mainTypes.Result, error) {
	if req.Id == 0 || req.Address == "" {
		return &mainTypes.Result{
			Code: 400,
			Msg:  "参数错误",
		}, nil
	}
	tx := global.GVA_DB.Begin()
	req.CreateTime = time.Now()
	err := tx.Create(&req).Error
	if err != nil {
		global.GVA_LOG.Error("服务器内部错误", zap.Any("err", err))
		tx.Rollback()
		return &mainTypes.Result{
			Code: 400,
			Msg:  "服务器内部错误",
		}, nil
	}
	wallet := types.Wallet{
		Uid:        req.Id,
		Currencyid: 1,
		Name:       "JHE",
		Balance:    0,
		Lock:       0,
	}

	err = tx.Create(&wallet).Error
	if err != nil {
		tx.Rollback()
		global.GVA_LOG.Error("服务器内部错误", zap.Any("err", err))
		return &mainTypes.Result{
			Code: 400,
			Msg:  "服务器内部错误",
		}, nil
	}
	tx.Commit()
	return &mainTypes.Result{
		Code: 200,
		Msg:  "新增用户成功",
	}, nil
}
func (h *HallUserLogic) AccountManager(req types.CustomerOperator) (*mainTypes.Result, error) {
	if req.Id == 0 || req.Num == 0 || req.Balance == 0 || req.Type == 0 {
		return &mainTypes.Result{
			Code: 400,
			Msg:  "参数错误",
		}, nil
	}
	req.CreateTime = time.Now()
	tx := global.GVA_DB.Begin()

	err := tx.Model(&req).Create(&req).Error
	if err != nil {
		global.GVA_LOG.Error("账户变动失败", zap.Any("err", err))
		tx.Rollback()
		return &mainTypes.Result{
			Code: 400,
			Msg:  "失败",
		}, nil
	}
	err = tx.Model(&types.Wallet{}).Where("uid = ?", req.Uid).Update("balance", req.Balance).Error
	if err != nil {
		global.GVA_LOG.Error("账户变动失败", zap.Any("err", err))
		tx.Rollback()
		return &mainTypes.Result{
			Code: 400,
			Msg:  "失败",
		}, nil
	}
	tx.Commit()
	return &mainTypes.Result{
		Code: 200,
		Msg:  "成功",
	}, nil
}
func (h *HallUserLogic) RewardCallBack(req types.CustomerOperator) (*mainTypes.Result, error) {
	tx := global.GVA_DB.Begin()

	err := tx.Model(&mainTypes.UserWithdrawl{}).Where("id = ?", req.Id).Update("status", 2).Update("create_time", time.Now()).Error
	if err != nil {
		global.GVA_LOG.Error("修改提币申请状态错误", zap.Any("err", err))
		tx.Rollback()
		return &mainTypes.Result{
			Code: 400,
			Msg:  "操作错误" + err.Error(),
		}, nil
	}
	var item = types.CustomerOperator{
		Uid:        req.Uid,
		Type:       2,
		Num:        req.Num,
		Balance:    req.Balance,
		CreateTime: time.Now(),
	}
	err = tx.Model(&types.CustomerOperator{}).Create(&item).Error
	if err != nil {
		global.GVA_LOG.Error("添加用户操作错误", zap.Any("err", err))
		tx.Rollback()
		return &mainTypes.Result{
			Code: 400,
			Msg:  "操作错误" + err.Error(),
		}, nil
	}
	err = global.GVA_DB.Model(&mainTypes.Wallet{}).Where("uid = ?", req.Uid).Update("balance", req.Balance).Error
	if err != nil {
		global.GVA_LOG.Error("修改钱包余额失败", zap.Any("err", err))
		tx.Rollback()
		return &mainTypes.Result{
			Code: 400,
			Msg:  "操作错误" + err.Error(),
		}, nil
	}
	tx.Commit()
	return &mainTypes.Result{
		Code: 200,
		Msg:  "操作成功",
	}, nil
}
func (h *HallUserLogic) GameRankConfig() (*mainTypes.Result, error) {
	var rankconfig []mainTypes.GameRankConfig
	day, week, month, _, err := model.GetGameConfig()
	if err != nil {
		global.GVA_LOG.Error("获取游戏排名奖励设置失败", zap.Any("err", err))
		return &mainTypes.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	if err := global.GVA_DB.Model(&rankconfig).Limit(30).Find(&rankconfig).Error; err != nil {
		global.GVA_LOG.Error("获取游戏排名奖励设置失败", zap.Any("err", err))
		return &mainTypes.Result{
			Code: 400,
			Msg:  "获取失败",
		}, nil
	}
	return &mainTypes.Result{
		Code: 200,
		Msg:  "获取成功",
		Data: types.Config{
			RankConfig: rankconfig,
			ProfitConfig: mainTypes.GameConfig{
				Day:   day,
				Week:  week,
				Month: month,
			},
		},
	}, nil
}
