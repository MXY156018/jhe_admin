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
}

func NewServiceContext(c config.Server) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		OperateRecord: middleware.OperationRecord,
		Jwt:           middleware.JWTAuth,
		Casbin:        middleware.CasbinHandler,
	}
}
