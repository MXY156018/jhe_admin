package checkhandler

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin, serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/check/passCheck",
					Handler: PassCheckHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/check/passCheckByIds",
					Handler: PassCheckByIdsHandler(serverCtx),
				},
			}...,
		),
	)
}
func RegisterHandlersAutocode1(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/check/getCheckList",
					Handler: GetCheckHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/check/getDailyReward",
					Handler: GetDailyRewardHandler(serverCtx),
				},
			}...,
		),
	)
}
