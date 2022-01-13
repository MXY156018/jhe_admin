package logic

import (
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
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

func (u *UserLogic) GetUserInfo() (*types.Result, error) {
	return &types.Result{}, nil
}
