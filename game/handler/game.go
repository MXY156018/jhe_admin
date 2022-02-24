/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 19:25:53
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-15 10:36:49
 */
package handler

import (
	"JHE_admin/internal/svc"
	"JHE_admin/utils"

	"JHE_admin/game/logic"
	"JHE_admin/game/types"

	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 冻结一定余额进行游戏
func FreezeAssetForGame(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GameFreezeAssetReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewGameLogic(r.Context(), ctx)
		resp, err := l.FreezeAssetForGame(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// 游戏结算
func GameSettle(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GameSettleReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewGameLogic(r.Context(), ctx)
		resp, err := l.Settle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
