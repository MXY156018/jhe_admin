package systemhandler

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
			{
				Method:  http.MethodPost,
				Path:    "/system/getServerInfo",
				Handler: GetServerInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/system/reloadSystem",
				Handler: ReloadSystemHandler(serverCtx),
			},
		},
	)
}
