package simpleuploadhandler

import (
	"JHE_admin/internal/logic/simpleupload"
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func SimpleUploaderUploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := simpleupload.NewSimpleUploadLogic(r.Context(), ctx)
		resp, err := l.SimpleUploaderUpload(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func CheckFileMd5Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := simpleupload.NewSimpleUploadLogic(r.Context(), ctx)
		resp, err := l.CheckFileMd5(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func MergeFileMd5Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := simpleupload.NewSimpleUploadLogic(r.Context(), ctx)
		resp, err := l.MergeFileMd5(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
