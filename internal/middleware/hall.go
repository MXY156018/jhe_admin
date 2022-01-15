package middleware

import (
	"JHE_admin/internal/types"
	"JHE_admin/web/hall/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func HallCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("uid")
		token := r.URL.Query().Get("token")
		if uid == "" {
			res := types.Result{Code: 401, Msg: "参数错误"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
		if token == "" {
			res := types.Result{Code: 401, Msg: "非法用户"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}

		covuid, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			res := types.Result{Code: 401, Msg: "参数错误"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
		isCheck := service.CheckHallUser(covuid, token)
		if isCheck {
			next(w, r)
		} else {
			res := types.Result{Code: 401, Msg: "非法用户"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}

	}
}
