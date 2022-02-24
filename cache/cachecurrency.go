/*
 * @Descripttion: 币种信息缓存
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:15:28
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 18:06:57
 */

package cache

import (
	"JHE_admin/table"

	"context"

	"gorm.io/gorm"
)

// 币种缓存
type CacheCurrency struct {
	isLoad   bool
	currency []table.Currency
}

// 加载参数
func (l *CacheCurrency) Load(db *gorm.DB) error {
	tx := db.WithContext(context.Background())
	err := tx.Find(&l.currency).Error
	if err != nil {
		return err
	}
	l.isLoad = true
	return nil
}

// 获取币种信息
func (l *CacheCurrency) GetCurrency(symbol string) *table.Currency {
	if !l.isLoad {
		return nil
	}
	for i := 0; i < len(l.currency); i++ {
		item := &l.currency[i]
		if item.Symbol == symbol {
			return item
		}
	}
	return nil
}
