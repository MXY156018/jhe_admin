/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-22 15:35:48
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 18:04:04
 */
package system

import (
	"JHE_admin/internal/logic/systemconfig"
	"JHE_admin/internal/svc"
	"JHE_admin/table"
	"JHE_admin/utils"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetSystemConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := systemconfig.NewSystemConfigLogic(r.Context(), ctx)
		resp, err := l.GetSystemConfig()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func SetSystemConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req []table.SysConfig
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := systemconfig.NewSystemConfigLogic(r.Context(), ctx)
		resp, err := l.SetSystemConfig(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
