package jwthandler

import (
	"JHE_admin/internal/logic/jwt"
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func JsonInBlacklistHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := jwt.NewJwtLogic(r.Context(), ctx)
		resp, err := l.JsonInBlacklist(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
