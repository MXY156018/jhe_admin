/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 23:13:46
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-20 15:56:24
 */
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
					Path:    "/check/backCheck",
					Handler: BackCheckHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/check/withdraw/fallback",
					Handler: WithdrawFallback(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/check/withdraw/markhandle",
					Handler: WithdrawFailManulHandle(serverCtx),
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
