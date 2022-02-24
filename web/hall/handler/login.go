/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 01:19:58
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 12:28:18
 */
package halluser

import (
	"JHE_admin/internal/svc"
	"JHE_admin/utils"
	"JHE_admin/web/hall/logic"
	"JHE_admin/web/hall/types"
	"net"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 账号密码登录
func AccountLogin(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountLoginReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		resp := l.AccountLogin(&req, ip)
		httpx.OkJson(w, resp)
	}
}
