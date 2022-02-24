/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:42:42
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 20:07:22
 */

package table

type Currency struct {
	// 币种
	Symbol string `gorm:"primaryKey" json:"symbol"`
	// 是否是体验币
	IsExp int8 `json:"isExp"`
	// 最小提币额度
	MinWithdraw float32 `json:"minWithdraw"`
	// 最大提币额度
	MaxWithdraw float32 `json:"maxWithdraw"`
	// 币的链地址
	Address string `json:"address"`
	// 精度
	Precision int8 `json:"precision"`
}

func (l *Currency) TableName() string {
	return "currency"
}
