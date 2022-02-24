/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 19:44:21
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 15:33:15
 */
package types

// 游戏结算条目
type GameFreezeAssetItem struct {
	// 用户uid
	UID int `json:"uid"`
	// 额度
	Amount float32 `json:"amount"`
}

// 冻结游戏额度请求
type GameFreezeAssetReq struct {
	// 币种
	Symbol string `json:"symbol"`
	// 是否是冻结
	IsFreeze bool `json:"isFreeze"`
	// 冻结条目
	Items []GameFreezeAssetItem `json:"items"`
}

// 冻结时返回
type GameFreezeResp struct {
	// 用户uid
	UID int `json:"uid"`
	// 用户余额
	Amount float32 `json:"amount"`
}

// 游戏结算条目
type GameSettleItem struct {
	// 用户uid
	UID int `json:"uid"`
	//赢分
	Win float32 `json:"win"`
	// 解冻 (押分)
	Unfreeze float32 `json:"unFreeze"`
	//抽水
	Rebate float32 `json:"rebate"`
}

// 获取用户信息请求
type GameSettleReq struct {
	// 游戏ID
	GameId int `json:"gameId"`
	// 币种
	Symbol string `json:"symbol"`
	// 结算条目
	Items []GameSettleItem `json:"items"`
}
