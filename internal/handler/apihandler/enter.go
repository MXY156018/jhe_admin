/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-24 17:34:21
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 17:45:01
 */
package apihandler

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
				Path:    "/api/createApi",
				Handler: CreateApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deleteApi",
				Handler: DeleteApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/getApiList",
				Handler: GetApiListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/getApiById",
				Handler: GetApiByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/updateApi",
				Handler: UpdateApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/getAllApis",
				Handler: GetAllApisHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deleteApisByIds",
				Handler: DeleteApisByIdsHandler(serverCtx),
			},
		},
	)
}
