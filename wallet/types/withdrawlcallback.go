/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 00:28:54
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 20:48:58
 */
package types

// 提币回调请求
type WithdrawlCallbackReq struct {
	// 未使用
	AppUid int `json:"appUid"`
	// 未使用
	Nonce string `json:"nonce"`
	//交易HASH
	Hash string `json:"hash,optional"`
	// 从那个地址转入
	From string `json:"from"`
	// 转到那个地址
	To string `json:"to"`
	// 时间戳
	Timestamp string `json:"timestamp,optional"`
	// 额度
	Quantity string `json:"quantity"`
	// 签名
	Sign string `json:"sign"`
	// 是否成功
	Success int8 `json:"success"`
	// 订单 ID
	OrderId string `json:"orderId"`
	// 错误码
	ErrCode int8 `json:"errCode,optional"`
	//币种
	Symbol string `json:"symbol"`
}
