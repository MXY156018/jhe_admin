/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-24 17:58:02
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 17:58:02
 */
package middleware

import (
	"net/http"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}
func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//跨域请求头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, User-Agent, x-token, token,x-user-id")
		next(w, r)
	}
}
func (m *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if r.Method != "OPTIONS" || origin == "" {
		if r.Response != nil {
			r.Response.StatusCode = http.StatusMethodNotAllowed
		}
		return
	}
	//跨域请求头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, User-Agent, x-token, token,x-user-id")
}
