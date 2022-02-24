/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:46:13
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 14:31:15
 */
package table

type UserRecharge struct {
	// 记录ID
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// 用户 UID
	UID int `json:"uid"`
	// 日期
	Date string `json:"date"`
	// 币种
	Symbol string `json:"symbol"`
	// 到账额度
	Amount float32 `json:"amount"`
	// 交易 hash
	TxHash string `json:"txHash"`
	// 原始额度
	RawAmount float32 `json:"rawAmount"`
	// 手续费
	Fee float32 `json:"fee"`
	// 充值前金额
	PreBalance float32 `json:"pre_balance"`
	// 充值后金额
	NowBalance float32 `json:"now_balance"`
}

func (l *UserRecharge) TableName() string {
	return "user_recharges"
}
