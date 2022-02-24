/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 03:03:50
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 14:51:09
 */
package table

type UserWithdrawl struct {
	// 订单ID
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// 用户UID
	UID int `json:"uid"`
	// 状态 0 请求  1 提币中 2 提币成功 3 提币失败  4 提币失败已经退回
	Status int8 `json:"status"`
	// 请求时间
	ReqDate string `json:"reqDate"`
	// 完成时间
	FinishDate string `json:"finishDate"`
	// 币种
	Symbol string `json:"symbol"`
	// 到账额度
	Amount float32 `json:"amount"`
	// 交易HASH
	TxHash string `json:"txHash"`
	// 原始额度
	RawAmount float32 `json:"rawAmount"`
	// 手续费
	Fee float32 `json:"fee"`
	// 提币到哪个地址
	To string `json:"to"`
	// 提币前金额
	PreBalance float32 `json:"pre_balance"`
	// 提币后金额
	NowBalance float32 `json:"now_balance"`
}

func (l *UserWithdrawl) TableName() string {
	return "user_withdrawls"
}

func (l *UserWithdrawl) GetSumWithdraw(date string) {

}
