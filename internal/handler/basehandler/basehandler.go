package basehandler

import (
	"JHE_admin/internal/logic/base"
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewCaptchaLogic(r.Context(), ctx)
		resp, err := l.Captcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
