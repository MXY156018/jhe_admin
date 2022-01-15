package halluser

import (
	"JHE_admin/internal/svc"
	"JHE_admin/utils"
	"JHE_admin/web/game/logic"
	"JHE_admin/web/game/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserHallLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HallUser
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewHallUserLogic(r.Context(), ctx)
		resp, err := l.UserLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
