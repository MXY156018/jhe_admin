/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-01-20 15:49:07
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-16 16:17:54
 */
package svc

import (
	"JHE_admin/internal/config"
	"JHE_admin/internal/middleware"

	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config        config.Server
	OperateRecord rest.Middleware
	Jwt           rest.Middleware
	Casbin        rest.Middleware
	Game          rest.Middleware
}

func NewServiceContext(c config.Server) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		OperateRecord: middleware.OperationRecord,
		Jwt:           middleware.JWTAuth,
		Casbin:        middleware.CasbinHandler,
		Game:          middleware.GameCheck,
	}
}
