/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 20:58:40
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-16 18:25:31
 */
package types

// 提币请求
type AssetWithdrawReq struct {
	// 币种
	Symbol string `json:"symbol"`
	// 额度
	Amount float32 `json:"amount"`
	// 提币到哪个地址
	To string `json:"to"`
}

// 批准提币
type AssetApproveWithdrawReq struct {
	// ID 订单ID
	ID int `json:"id"`
}

// 回退 用于提币失败的请求
type AssetFallbackWithdrawReq struct {
	// ID 订单ID
	ID int `json:"id"`
}
