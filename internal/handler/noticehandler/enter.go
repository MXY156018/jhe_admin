package noticehandler

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin, serverCtx.OperateRecord},
			//[]rest.Middleware{serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/notice/createNotice",
					Handler: CreateNoticeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/fleshNotice",
					Handler: FleshNoticeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/updateNotice",
					Handler: UpdateNoticeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/deleteNoticeByIds",
					Handler: DeleteNoticeByIdsHandler(serverCtx),
				},
			}...,
		),
	)
}
func RegisterHandlersAutocode1(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin},
			//[]rest.Middleware{serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/notice/getNoticeList",
					Handler: NoticeListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/getNoticeById",
					Handler: GetNoticeHandler(serverCtx),
				},
			}...,
		),
	)
}
