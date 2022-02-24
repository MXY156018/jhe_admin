/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-22 15:35:16
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 16:38:59
 */
package system

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
				Path:    "/system/getSystemConfig",
				Handler: GetSystemConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/system/setSystemConfig",
				Handler: SetSystemConfigHandler(serverCtx),
			},
		},
	)
}
