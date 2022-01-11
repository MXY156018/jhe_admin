package billhandler

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
					Path:    "/bill/getBill",
					Handler: BillHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/bill/getGameBill",
					Handler: GameBillHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/bill/getDailyBill",
					Handler: DailyBillHandler(serverCtx),
				},
			}...,
		),
	)
}
