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

func (c *CustomerLogic) GetCustomerList(req types.CustomerList) (*types.Result, error) {
	var list []types.CustomerList
	var total int64
	list, total, err := model.GetCustomerList(req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "获取失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: &types.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}, nil
}

func (c *CustomerLogic) ChangeCustomerStatus(req types.Customers) (*types.Result, error) {
	msg := ""
	if req.Status == 0 {
		msg = "禁用"
	} else {
		msg = "解禁"
	}
	if err := global.GVA_DB.Model(&types.Customers{}).Where("id = ?", req.Id).Update("status", req.Status).Error; err != nil {

		return &types.Result{
			Code: 7,
			Msg:  msg + "失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  msg + "成功",
	}, nil
}
func (c *CustomerLogic) GetCustomerById(req types.Customers) (*types.Result, error) {
	user, err := model.GetCustomerById(req)
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "获取失败" + err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: user,
	}, nil
}
func (c *CustomerLogic) DeleteCustomer(req types.Customers) (*types.Result, error) {
	if req.Id <= 0 {
		return &types.Result{
			Code: 7,
			Msg:  "参数错误",
		}, nil
	}
	err := global.GVA_DB.Delete(&types.Customers{}, req.Id).Error
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "删除失败" + err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
