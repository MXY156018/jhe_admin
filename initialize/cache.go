/*
 * @Descripttion: 初始化缓存
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 17:05:07
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-16 17:07:49
 */

package initialize

import (
	"JHE_admin/cache"
	"JHE_admin/global"
)

// 初始化缓存
func InitCache() error {
	global.GVA_CacheSysConfig = &cache.CacheSysConfig{}
	global.GVA_CacheCurrency = &cache.CacheCurrency{}

	err := global.GVA_CacheSysConfig.Load(global.GVA_DB)
	if err != nil {
		return err
	}
	err = global.GVA_CacheCurrency.Load(global.GVA_DB)
	if err != nil {
		return err
	}
	return nil
}
