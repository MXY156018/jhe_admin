/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 19:21:11
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-17 20:57:31
 */
// 用户鉴权及用户信息接口
package handler

import (
	"JHE_admin/internal/svc"
	"JHE_admin/utils"

	"JHE_admin/game/logic"
	"JHE_admin/game/types"

	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 根据 JWT token获取用户信息
func CheckJwt(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckJwtReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.CheckJwt(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// 根据 JWT token获取用户信息
func GetUserInfo(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserInfoByToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// 获取用户特定资产
func GetUserAsset(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserAssetReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserAsset(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// 获取用户所有资产
func GetUserAssets(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserAssetsReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserAssets(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
