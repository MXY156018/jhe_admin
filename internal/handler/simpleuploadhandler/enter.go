package simpleuploadhandler

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/simpleUploader/upload",
					Handler: SimpleUploaderUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/simpleUploader/checkFileMd5",
					Handler: CheckFileMd5Handler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/simpleUploader/mergeFileMd5",
					Handler: MergeFileMd5Handler(serverCtx),
				},
			}...,
		),
	)
}
