/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 23:47:55
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 18:27:20
 */

package table

type UserBlockchainAccount struct {
	// 用户UID
	UID int `gorm:"primaryKey" json:"uid"`
	// 充值地址
	Address string `json:"address"`
}

func (l *UserBlockchainAccount) TableName() string {
	return "user_block_chain_accounts"
}
