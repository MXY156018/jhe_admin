/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 00:44:26
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-17 17:13:08
 */

package handler

import (
	"JHE_admin/internal/svc"
	"JHE_admin/wallet/middleware"

	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func Register(engine *rest.Server, serverCtx *svc.ServiceContext) {
	auth := middleware.NewWhiteIpAuth()
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{auth.Handler},
			[]rest.Route{
				{
					//充值回调
					Method:  http.MethodPost,
					Path:    "/api/wallet/recharge",
					Handler: RechargeCallback(serverCtx),
				},
				{
					//提币回调
					Method:  http.MethodPost,
					Path:    "/api/wallet/withdrawl",
					Handler: WithdrawlCallback(serverCtx),
				},
			}...,
		),
	)
}
