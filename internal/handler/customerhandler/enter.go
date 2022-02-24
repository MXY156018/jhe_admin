/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 15:13:40
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-23 14:36:01
 */
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
					//获取用户列表
					Method:  http.MethodPost,
					Path:    "/customer/getCustomerList",
					Handler: CustomerListHandler(serverCtx),
				},
				{
					//改变用户状态
					Method:  http.MethodPost,
					Path:    "/customer/changeCustomerStatus",
					Handler: CustomerStatusHandler(serverCtx),
				},
				{
					//按id获取用户信息
					Method:  http.MethodPost,
					Path:    "/customer/getCustomerById",
					Handler: GetCustomerByIdHandler(serverCtx),
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
func RegisterHandlersAutocode1(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/customer/getGameRecord",
					Handler: GetCustomerGameRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customer/getOperator",
					Handler: GetCustomerOperatorHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/home/homeData",
					Handler: HomeDataHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customer/getUserWallet",
					Handler: GetUserWallet(serverCtx),
				},
			}...,
		),
	)
}
