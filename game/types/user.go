/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 19:12:29
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-17 20:41:25
 */
package types

// 获取用户信息请求
type CheckJwtReq struct {
	// JWT token
	Token string `json:"token"`
}

// 获取用户信息请求
type GetUserInfoReq struct {
	// JWT token
	UID int `json:"uid"`
}

// 获取用户信息回复
type GetUserInfoResp struct {
}

// 获取用户余额请求
type GetUserAssetReq struct {
	// 币种
	Symbol string `json:"symbol"`
	// 用户UID
	UID int `json:"uid"`
}

// 获取用户资产
type GetUserAssetsReq struct {
	// 用户UID
	UID int `json:"uid"`
}

type AssetUpdateItem struct {
	// 用户UID
	UID int `json:"uid"`
	// 币种
	Symbol string `json:"symbol"`
	// 变更的额度
	Amount float32 `json:"amount"`
	// 冻结变更的额度
	FreezeAmount float32 `json:"freezeAmount"`
}

// 更改用户余额请求
type UpdateUserAssetReq struct {
	// 币种
	Assets []AssetUpdateItem `json:"assets"`
}
