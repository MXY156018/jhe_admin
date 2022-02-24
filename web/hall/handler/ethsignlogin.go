/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 18:25:15
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 14:20:19
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

// 通过ETH签名方式登录
func EthSignLogin(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EthSignLoginReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEthSignLoginLogic(r.Context(), ctx)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		resp := l.Login(&req, ip)
		httpx.OkJson(w, resp)
	}
}
