package halluser

import (
	"JHE_admin/internal/svc"
	"JHE_admin/web/hall/logic"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
