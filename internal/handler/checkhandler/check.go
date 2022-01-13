package checkhandler

import (
	"JHE_admin/internal/logic/check"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetCheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RewardReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp, err := l.GetCheckList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func PassCheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Reward
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp, err := l.PassCheck(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func PassCheckByIdsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Ids
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp, err := l.PassCheckByIds(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetDailyRewardHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err :=httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp, err := l.GetDailyReward()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
