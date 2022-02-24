/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-10 11:03:10
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 18:02:03
 */
package halluser

import (
	"JHE_admin/internal/svc"
	"JHE_admin/web/hall/middleware"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{middleware.JWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/wallet",
					Handler: GetWalletHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/userfeedback",
					Handler: FeedBackHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/hall/getNotice",
					Handler: Notice(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getFeedBack",
					Handler: FeedBack(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/readFeedBack",
					Handler: ReadFeedBack(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/hall/getRankList",
					Handler: GetRankList(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getRechargeList",
					Handler: GetRechargeList(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getRewardList",
					Handler: GetRewardList(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getProfitList",
					Handler: GetProfitList(serverCtx),
				},

				{
					Method:  http.MethodPost,
					Path:    "/hall/userWithdraw",
					Handler: UserWithdraw(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getConfig",
					Handler: GetConfig(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getProfit",
					Handler: GetProfit(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getCurrency",
					Handler: GetCurrency(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getBanner",
					Handler: GetBanner(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/drawProfit",
					Handler: DrawProfit(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/getVipConfig",
					Handler: VipConfig(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hall/buyVIP",
					Handler: BuyVip(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/game/gameRank",
					Handler: GetInGameRank(serverCtx),
				},
			}...,
		),
	)
	engine.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodPost,
				Path:   "/user/halllogin",
				// Handler: UserHallLoginHandler(serverCtx),
				Handler: AccountLogin(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/ethlogin",
				Handler: EthSignLogin(serverCtx),
			},
		},
		// rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
