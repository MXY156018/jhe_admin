/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 00:34:45
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 20:48:52
 */
package types

// 充值请求
type RechargeCallbackReq struct {
	// 未使用
	AppUid int `json:"appUid"`
	// 未使用
	Nonce string `json:"nonce"`
	//币种
	Symbol string `json:"symbol"`
	//交易HASH
	Hash string `json:"hash"`
	// 从那个地址转入
	From string `json:"from"`
	// 转到那个地址
	To string `json:"to"`
	// 时间戳
	Timestamp string `json:"timestamp"`
	// 额度
	Quantity string `json:"quantity"`
	// 签名
	Sign string `json:"sign"`
}
