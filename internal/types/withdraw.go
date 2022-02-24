/*
 * @Descripttion: 提币管理
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 22:03:26
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 23:14:49
 */

package types

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

// 提币失败，已经手动处理
type AssetWithdrawFailManulHandleReq struct {
	// ID 订单ID
	ID int `json:"id"`
}
