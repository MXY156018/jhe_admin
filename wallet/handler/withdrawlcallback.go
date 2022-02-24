/*
 * @Descripttion: 提币回调
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 00:28:07
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-15 00:44:08
 */
package handler

import (
	"JHE_admin/internal/svc"
	"JHE_admin/utils"

	"JHE_admin/wallet/logic"
	"JHE_admin/wallet/types"

	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 充值回调
func WithdrawlCallback(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawlCallbackReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewWithdrawlCallbackLogic(r.Context(), ctx)
		resp, err := l.Callback(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
