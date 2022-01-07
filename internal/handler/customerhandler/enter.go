package customerhandler

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
					Path:    "/customer/getCustomerList",
					Handler: CustomerListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customer/changeCustomerStatus",
					Handler: CustomerStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customer/getCustomerById",
					Handler: GetCustomerByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customer/deleteCustomerById",
					Handler: DeleteCustomerByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customer/getSubordinate",
					Handler: GetSubordinateHandler(serverCtx),
				},
			}...,
		),
	)
}
