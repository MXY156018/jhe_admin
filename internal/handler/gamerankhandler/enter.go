package gamerankhandler

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
					Path:    "/gamerank/setConfig",
					Handler: SetGameRankConfigHandler(serverCtx),
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
					Method:  http.MethodGet,
					Path:    "/gamerank/getConfig",
					Handler: GetGameRankConfigHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/gamerank/getRank",
					Handler: GetGameRankHandler(serverCtx),
				},
			}...,
		),
	)
}
