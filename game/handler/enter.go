/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 10:05:44
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-17 10:38:18
 */
package handler

// 游戏 API 注册入口

import (
	"JHE_admin/game/middleware"
	"JHE_admin/internal/svc"
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
					//获取用户信息
					Method:  http.MethodPost,
					Path:    "/api/game/checkjwt",
					Handler: CheckJwt(serverCtx),
				},
				{
					//获取用户信息
					Method:  http.MethodPost,
					Path:    "/api/game/userinfo",
					Handler: GetUserInfo(serverCtx),
				},
				{
					//获取用户固定资产
					Method:  http.MethodPost,
					Path:    "/api/game/userasset",
					Handler: GetUserAsset(serverCtx),
				},
				{
					//获取用户所有资产
					Method:  http.MethodPost,
					Path:    "/api/game/userassets",
					Handler: GetUserAssets(serverCtx),
				},

				{
					//冻结余额游戏
					Method:  http.MethodPost,
					Path:    "/api/game/freezeasset",
					Handler: FreezeAssetForGame(serverCtx),
				},
				{
					//游戏结算
					Method:  http.MethodPost,
					Path:    "/api/game/settle",
					Handler: GameSettle(serverCtx),
				},
			}...,
		),
	)
}
