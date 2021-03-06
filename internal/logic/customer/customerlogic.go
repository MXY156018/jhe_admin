package customer

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type CustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) CustomerLogic {
	return CustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *CustomerLogic) GetCustomerList(req types.CustimerSearch) (*types.Result, error) {
	var total int64
	list, total, err := model.GetCustomerList(req)
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

func (c *CustomerLogic) ChangeCustomerStatus(req types.UserReq) (*types.Result, error) {
	msg := ""
	if req.Status == 0 {
		msg = "禁用"
	} else {
		msg = "解禁"
	}
	if err := global.GVA_DB.Table("users").Where("uid = ?", req.Uid).Update("status", req.Status).Error; err != nil {

		return &types.Result{
			Code: 7,
			Msg:  msg + "失敗",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  msg + "成功",
	}, nil
}
func (c *CustomerLogic) GetCustomerById(req types.UserDetail) (*types.Result, error) {
	user, err := model.GetCustomerById(req)
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗" + err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "獲取成功",
		Data: user,
	}, nil
}
func (c *CustomerLogic) GetWallet(req types.Wallet) (*types.Result, error) {

	var wallet []types.Wallet
	db := global.GVA_DB.WithContext(context.Background())
	err := db.Model(&wallet).Where("uid = ?", req.Uid).Find(&wallet).Error
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗" + err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "獲取成功",
		Data: wallet,
	}, nil
}

// func (c *CustomerLogic) DeleteCustomer(req types.Customers) (*types.Result, error) {
// 	if req.Id <= 0 {
// 		return &types.Result{
// 			Code: 7,
// 			Msg:  "參數錯誤",
// 		}, nil
// 	}
// 	err := global.GVA_DB.Delete(&types.Customers{}, req.Id).Error
// 	if err != nil {
// 		return &types.Result{
// 			Code: 7,
// 			Msg:  "刪除失敗" + err.Error(),
// 		}, nil
// 	}
// 	return &types.Result{
// 		Code: 0,
// 		Msg:  "刪除成功",
// 	}, nil
// }
func (c *CustomerLogic) GetSubordinate(req types.CustimerSearch) (*types.Result, error) {
	if req.Uid <= 0 {
		// global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "參數錯誤",
		}, nil
	}
	list, total, err := model.GetSubordinateModel(req)
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
func (c *CustomerLogic) GetCustomerGameRerord(req types.GameRecordList) (*types.Result, error) {
	list, total, err := model.GetCustomerGameRecordModel(req)
	if err != nil {
		global.GVA_LOG.Error("獲取失敗!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取成功",
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
func (c *CustomerLogic) GetCustomerOperator(req types.OperateRecord) (*types.Result, error) {
	list, total, err := model.GetCustomerOperatorModel(req)
	if err != nil {
		global.GVA_LOG.Error("獲取事變!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取成功",
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
func (c *CustomerLogic) GetHomeData() (*types.Result, error) {

	var userNum int64
	err := global.GVA_DB.Model(&types.User{}).Where("is_bot = 0").Count(&userNum).Error
	if err != nil {
		global.GVA_LOG.Error("服務器內部錯誤", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	platform, err := model.GetSumPlatformProfit()
	if err != nil {
		global.GVA_LOG.Error("服務器內部錯誤", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	rechargeNum, err := model.GetSum(1)
	if err != nil {
		global.GVA_LOG.Error("服務器內部錯誤", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	rewardNum, err := model.GetSum(2)
	if err != nil {
		global.GVA_LOG.Error("服務器內部錯誤", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	vipprofitNum, err := model.GetSum(3)
	if err != nil {
		global.GVA_LOG.Error("服務器內部錯誤", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}
	gameprogitNum, err := model.GetSum(4)
	if err != nil {
		global.GVA_LOG.Error("服務器內部錯誤", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "獲取失敗",
		}, nil
	}

	return &types.Result{
		Code: 0,
		Msg:  "獲取成功",
		Data: &types.HomeData{
			UserNum:        userNum,
			PlatformProfit: platform,
			RechargeNum:    rechargeNum,
			RewardNum:      rewardNum,
			VipbalanceNum:  vipprofitNum,
			GameablanceNum: gameprogitNum,
		},
	}, nil
}
