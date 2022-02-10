package main

import (
	"JHE_admin/core"
	"JHE_admin/global"
	"JHE_admin/initialize"
	"JHE_admin/internal/config"
	"JHE_admin/internal/handler"
	"JHE_admin/internal/svc"
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/jhe-admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Server
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	global.GVA_VP = core.Viper()               // 初始化Viper
	global.GVA_LOG = core.Zap(c)               // 初始化zap日志库
	global.GVA_DB = initialize.Gorm(c.Mysql)   // gorm连接数据库
	global.GVA_DB2 = initialize.Gorm(c.Mysql2) // gorm连接数据库
	global.GVA_CONFIG = c
	initialize.Timer()

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	//server.Use(middleware.JWTAuth)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
