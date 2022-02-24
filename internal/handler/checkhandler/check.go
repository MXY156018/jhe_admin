/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 13:57:23
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 23:12:56
 */
package checkhandler

import (
	"JHE_admin/internal/logic/check"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/utils"

	// "JHE_admin/web/hall/logic"
	// hallTypes "JHE_admin/web/hall/types"
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

// 批准提币
func PassCheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetApproveWithdrawReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp := l.ApproveWithdraw(&req)
		httpx.OkJson(w, resp)
	}
}
func BackCheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetApproveWithdrawReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp := l.BackWithdraw(&req)
		httpx.OkJson(w, resp)
	}
}

// 退回
func WithdrawFallback(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetFallbackWithdrawReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp := l.WithdrawFallBack(&req)
		httpx.OkJson(w, resp)
	}
}

// 提币失败，已经手动处理
func WithdrawFailManulHandle(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetWithdrawFailManulHandleReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := check.NewCheckLogic(r.Context(), ctx)
		resp := l.WithdrawFailManulHandle(&req)
		httpx.OkJson(w, resp)
	}
}

//未使用
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
		if err := httpx.Parse(r, &req); err != nil {
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
