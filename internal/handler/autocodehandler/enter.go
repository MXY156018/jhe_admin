package autocodehandler

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/autoCode/delSysHistory",
					Handler: DelSysHistoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/autoCode/getMeta",
					Handler: GetMetaHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/autoCode/getSysHistory",
					Handler: GetSysHistoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/autoCode/rollback",
					Handler: RollBackHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/autoCode/preview",
					Handler: PreviewTempHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/autoCode/createTemp",
					Handler: CreateTempHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/autoCode/getTables",
					Handler: GetTablesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/autoCode/getDB",
					Handler: GetDBHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/autoCode/getColumn",
					Handler: GetColumnHandler(serverCtx),
				},
			}...,
		),
	)
}
