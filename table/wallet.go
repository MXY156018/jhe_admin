/*
 * @Descripttion: 用户钱包
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-17 21:35:16
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 02:40:09
 */
package table

type Wallet struct {
	UID            int     `gorm:"primaryKey" json:"uid"`
	Symbol         string  `gorm:"primaryKey" json:"symbol"`
	Balance        float32 `json:"balance"`
	GameFrozen     float32 `json:"gameFronze"`
	WithdrawFrozen float32 `json:"withdrawFrozen"`
}

func (l *Wallet) TableName() string {
	return "wallets"
}
