package casbinhandler

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
					Path:    "/casbin/updateCasbin",
					Handler: UpdateCasbinHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/casbin/getPolicyPathByAuthorityId",
					Handler: GetPolicyPathByAuthorityIdHandler(serverCtx),
				},
			}...,
		),
	)
}
