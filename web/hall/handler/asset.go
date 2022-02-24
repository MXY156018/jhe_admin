/*
 * @Descripttion: 用户资产相关接口
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 20:58:06
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 22:21:36
 */

package halluser

import (
	"JHE_admin/internal/svc"
	"JHE_admin/utils"
	"JHE_admin/web/hall/logic"
	"JHE_admin/web/hall/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 提币请求
func Withdraw(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetWithdrawReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAssetLogic(r.Context(), ctx)
		resp, err := l.Withdraw(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// // 审批提币请求
// func ApproveWithdraw(ctx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req types.AssetApproveWithdrawReq
// 		if err := utils.Bind(r, &req); err != nil {
// 			httpx.Error(w, err)
// 			return
// 		}

// 		l := logic.NewAssetLogic(r.Context(), ctx)
// 		resp, err := l.ApproveWithdraw(&req)
// 		if err != nil {
// 			httpx.Error(w, err)
// 		} else {
// 			httpx.OkJson(w, resp)
// 		}
// 	}
// }

// // 提币失败退回
// func WithdrawFailFallback(ctx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req types.AssetFallbackWithdrawReq
// 		if err := utils.Bind(r, &req); err != nil {
// 			httpx.Error(w, err)
// 			return
// 		}

// 		l := logic.NewAssetLogic(r.Context(), ctx)
// 		resp, err := l.FallBack(&req)
// 		if err != nil {
// 			httpx.Error(w, err)
// 		} else {
// 			httpx.OkJson(w, resp)
// 		}
// 	}
// }

// 获取资产记录
func GetAssetHistory(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
