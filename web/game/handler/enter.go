package halluser

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Game},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/halllogin",
					Handler: UserHallLoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/userRegister",
					Handler: UserRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/userAccountManager",
					Handler: UserAccountManagerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/userRewardCallBack",
					Handler: UserRewardCallBackHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getGameRankConfig",
					Handler: GetGameRankConfigHandler(serverCtx),
				},
			}...,
		),
	)
}
