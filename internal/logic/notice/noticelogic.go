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
