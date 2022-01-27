package halluser

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/userinfo",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/userfeedback",
				Handler: FeedBackHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/test/getUserTree",
				Handler: UserTreeHandler(serverCtx),
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
				Path:    "/hall/getVipProfit",
				Handler: GetVipProfitList(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/hall/getGameProfit",
				Handler: GetGameProfitList(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
