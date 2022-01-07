package jwt

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"context"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type JwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) JwtLogic {
	return JwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//JsonInBlacklist jwt加入黑名单
func (a *JwtLogic) JsonInBlacklist(r *http.Request) (*types.Result, error) {
	token := r.Header.Get("x-token")
	jwt := types.JwtBlacklist{Jwt: token}
	if err := model.JwtServiceApp.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "jwt作废失败"}, nil
	}
	return &types.Result{Code: 7, Msg: "jwt作废成功"}, nil
}
