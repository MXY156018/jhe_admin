/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 23:25:09
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-15 15:08:38
 */
package middleware

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"context"
	"net"

	"net/http"
	"sync"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type ip struct {
	IP string
}

// 游戏 IP 通过白名单验证
type WhiteIpAuth struct {
	// 是否已经加载白名单
	isLoad bool
	// IP 白名单
	whiteIps []string
	// 锁
	lock sync.Mutex
}

func NewWhiteIpAuth() *WhiteIpAuth {
	return &WhiteIpAuth{}
}
func (m *WhiteIpAuth) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 加载 IP 白名单
		m.lock.Lock()
		if !m.isLoad {
			ips := []ip{}
			db := global.GVA_DB.WithContext(context.Background())
			db = db.Table("ip_white_list").Select("ip").Where("api_type=1").Find(&ips)
			if db.Error == nil {
				m.isLoad = true
				for i := 0; i < len(ips); i++ {
					m.whiteIps = append(m.whiteIps, ips[i].IP)
				}
			}
		}
		m.lock.Unlock()

		isFind := false
		addr, _, _ := net.SplitHostPort(r.RemoteAddr)
		for i := 0; i < len(m.whiteIps); i++ {
			if m.whiteIps[i] == addr {
				isFind = true
				break
			}
		}
		if isFind {
			next(w, r)
			return
		}

		httpx.OkJson(w, &types.Result{
			Code: 1,
			Msg:  "IP地址不在白名单中",
		})
	}
}
