package noticelogic

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

type NoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) NoticeLogic {
	return NoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (n *NoticeLogic) GetNoticeList(req types.NoticePage) (*types.Result, error) {

	total, list, err := model.NoticeList(req)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Data: &types.PageResult{
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
		Data: &types.PageResult{
			Total:    total,
			List:     list,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Msg: "获取成功",
	}, nil
}
func (n *NoticeLogic) CreateNotice(req types.Notice) (*types.Result, error) {
	var notice types.Notice
	if req.Title == "" {
		return &types.Result{
			Code: 7,
			Msg:  "请输入标题",
		}, nil
	}
	if req.Content == "" {
		return &types.Result{
			Code: 7,
			Msg:  "请输入内容",
		}, nil
	}

	notice = types.Notice{
		Title:      req.Title,
		Content:    req.Content,
		CreateTime: time.Now(),
	}
	if err := global.GVA_DB.Model(&notice).Create(&notice).Error; err != nil {
		global.GVA_LOG.Error("创建公告失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "创建成功",
	}, nil
}
func (n *NoticeLogic) GetNotice(req types.Notice) (*types.Result, error) {
	if req.Id <= 0 {
		return &types.Result{
			Code: 7,
			Msg:  "参数错误",
		}, nil
	}
	var notice types.Notice
	if err := global.GVA_DB.Model(&notice).Where("id = ?", req.Id).Find(&notice).Error; err != nil {
		global.GVA_LOG.Error("获取公告详情失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "获取公告详情失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "获取公告成功",
		Data: notice,
	}, nil
}
func (n *NoticeLogic) FleshNotice(req types.Notice) (*types.Result, error) {
	if req.Id <= 0 {
		return &types.Result{
			Code: 7,
			Msg:  "参数错误",
		}, nil
	}
	if err := global.GVA_DB.Model(&types.Notice{}).Where("id = ?", req.Id).Update("create_time", time.Now()).Error; err != nil {
		global.GVA_LOG.Error("刷新公告失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "刷新公告失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "刷新公告成功",
	}, nil
}
func (n *NoticeLogic) UpdateNotice(req types.Notice) (*types.Result, error) {
	if req.Id <= 0 {
		return &types.Result{
			Code: 7,
			Msg:  "参数错误",
		}, nil
	}
	req.CreateTime = time.Now()
	if err := global.GVA_DB.Model(&types.Notice{}).Where("id = ?", req.Id).Save(&req).Error; err != nil {
		global.GVA_LOG.Error("更新公告失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "更新公告失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "更新公告成功",
	}, nil
}
func (n *NoticeLogic) DeleteNoticeByIds(req types.Ids) (*types.Result, error) {
	if len(req.Id) <= 0 {
		return &types.Result{
			Code: 7,
			Msg:  "参数错误",
		}, nil
	}

	if err := global.GVA_DB.Model(&types.Notice{}).Where("id in ?", req.Id).Delete(&types.Notice{}).Error; err != nil {
		global.GVA_LOG.Error("批量删除公告失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "批量删除公告失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "批量删除公告成功",
	}, nil
}
