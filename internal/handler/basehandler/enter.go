/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-24 17:33:53
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 18:00:07
 */
package basehandler

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
				Path:    "/base/captcha",
				Handler: CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/base/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/jwt/jsonInBlacklist",
				Handler: JsonInBlacklistHandler(serverCtx),
			},
		},
	)
}
