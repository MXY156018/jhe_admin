package noticehandler

import (
	noticelogic "JHE_admin/internal/logic/notice"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func NoticeListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NoticePage
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := noticelogic.NewNoticeLogic(r.Context(), ctx)
		resp, err := l.GetNoticeList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func CreateNoticeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Notice
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := noticelogic.NewNoticeLogic(r.Context(), ctx)
		resp, err := l.CreateNotice(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
