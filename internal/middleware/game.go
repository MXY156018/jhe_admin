package middleware

import (
	"net/http"
)

func GameCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// isCheck := service.CheckHallUser(covuid, token)
		// if isCheck {
		// 	next(w, r)
		// } else {
		// 	res := types.Result{Code: 401, Msg: "非法用户"}
		// 	resp, _ := json.Marshal(res)
		// 	w.Write(resp)
		// 	return
		// }
		// res := types.Result{Code: 401, Msg: r.RemoteAddr}
		// resp, _ := json.Marshal(res)
		// w.Write(resp)
		// return
		next(w, r)
	}
}
