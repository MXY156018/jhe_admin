/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 11:28:59
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-23 11:52:45
 */
package middleware

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"JHE_admin/model"

	"encoding/json"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 拦截器
func CasbinHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := r.Header.Get("claims")
		var waitUse types.CustomClaims
		err := json.Unmarshal([]byte(str), &waitUse)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 获取请求的URI
		obj := r.RequestURI
		// 获取请求方法
		act := r.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := model.CasbinServiceApp.Casbin()

		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" || success {
			next(w, r)
		} else {
			res := types.Result{Code: 7, Msg: "權限不足符合"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
	}
}
