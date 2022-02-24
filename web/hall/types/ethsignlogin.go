/*
 * @Descripttion: ETH 钱包签名登录请求
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 18:25:37
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-16 22:55:54
 */

package types

// ETH签名登录请求参数
type EthSignLoginReq struct {
	// 钱包地址
	Address string `json:"address"`
	// 时间戳
	Timestamp int `json:"timestamp"`
	//签名消息
	Message string `json:"message"`
	// 消息的 SHA3 签名
	Sha3Message string `json:"sha3Message"`
	// 签名
	Sign string `json:"sign"`
	// 上级 可选
	Parent int `json:"parent,optional"`
}
