/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-17 15:37:19
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-17 17:34:25
 */
package middleware

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 我们这里jwt鉴权取头部信息 authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := r.Header.Get("authorization")
		if token == "" {
			res := types.Result{Code: 401, Msg: "未登錄或非法訪問"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
		}
		decodeToken, err := utils.DecodeGameJwtToken(token, global.GVA_CONFIG.Auth.AccessSecret)
		if err != nil {
			res := types.Result{Code: 402, Msg: "无效token" + err.Error()}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
		if decodeToken.Expire < int(time.Now().Unix()) {
			res := types.Result{Code: 403, Msg: "授權已過期,請重新登錄"}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "uid", decodeToken.UserID))
		next(w, r)
		//return
	}
}
