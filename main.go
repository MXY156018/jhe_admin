/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 17:08:55
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-23 12:09:41
 */
package main

import (
	"JHE_admin/core"
	"JHE_admin/global"
	"JHE_admin/initialize"
	"JHE_admin/internal/config"
	"JHE_admin/internal/handler"
	"JHE_admin/internal/middleware"
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
	corsmd := middleware.NewCorsMiddleware()
	cors := rest.WithNotAllowedHandler(corsmd)
	server := rest.MustNewServer(c.RestConf, cors)
	server.Use(corsmd.Handle)
	defer server.Stop()
	global.GVA_VP = core.Viper()       // 初始化Viper
	global.GVA_LOG = core.Zap(c)       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm(c) // gorm连接数据库
	global.GVA_CONFIG = c
	initialize.Timer()
	go initialize.NewCorn()
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 要放在 gorm 初始化之后
	err := initialize.InitCache()
	if err != nil {
		fmt.Printf("初始化缓存错误 %+v\n", err)
	}
	// times := "2020-01-11 15:04:05"
	// s, _ := dateparse.ParseAny(times)
	// fmt.Println(s)
	// str := utils.GetPreDate(time.Now())
	// fmt.Println(str)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
