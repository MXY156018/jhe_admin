/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 17:04:27
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 17:57:05
 */
package global

import (
	"JHE_admin/cache"
	"JHE_admin/utils/timer"

	"github.com/go-redis/redis"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"JHE_admin/internal/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}
	// 系统配置缓存
	GVA_CacheSysConfig *cache.CacheSysConfig
	// 币种缓存
	GVA_CacheCurrency *cache.CacheCurrency
)
