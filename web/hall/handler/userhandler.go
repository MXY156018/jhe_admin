/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-01-15 17:05:03
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 15:36:47
 */
package halluser

import (
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"JHE_admin/web/hall/logic"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetWalletHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserWallet()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetInGameRank(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetInGameRank()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetInGameRecord(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetInGameRecord(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
