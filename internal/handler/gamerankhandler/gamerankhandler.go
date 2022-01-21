package gamerankhandler

import (
	"JHE_admin/internal/logic/gamerank"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetGameRankConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := gamerank.NewGameRankLogic(r.Context(), ctx)
		resp, err := l.GetGameRankConfig()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func SetGameRankConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GameConfig
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := gamerank.NewGameRankLogic(r.Context(), ctx)
		resp, err := l.SetGameRankConfig(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetGameRankHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GameRankSearch
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := gamerank.NewGameRankLogic(r.Context(), ctx)
		resp, err := l.GetGameRank(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
