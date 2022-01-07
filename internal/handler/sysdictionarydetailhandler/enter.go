package sysdictionarydetailhandler

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
					Path:    "/sysDictionaryDetail/createSysDictionaryDetail",
					Handler: CreateSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/sysDictionaryDetail/deleteSysDictionaryDetail",
					Handler: DeleteSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/sysDictionaryDetail/updateSysDictionaryDetail",
					Handler: UpdateSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/sysDictionaryDetail/findSysDictionaryDetail",
					Handler: FindSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/sysDictionaryDetail/getSysDictionaryDetailList",
					Handler: GetSysDictionaryDetailListHandler(serverCtx),
				},
			}...,
		),
	)
}
