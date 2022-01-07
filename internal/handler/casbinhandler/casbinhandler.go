package casbinhandler

import (
	"JHE_admin/internal/logic/casbin"
	"JHE_admin/internal/types"
	"JHE_admin/utils"

	"github.com/tal-tech/go-zero/rest/httpx"

	"JHE_admin/internal/svc"
	"net/http"
)

func UpdateCasbinHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CasbinInReceive
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := casbin.NewCasbinLogic(r.Context(), ctx)
		resp, err := l.UpdateCasbin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetPolicyPathByAuthorityIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CasbinSearch
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := casbin.NewCasbinLogic(r.Context(), ctx)
		resp, err := l.GetPolicyPathByAuthorityId(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
