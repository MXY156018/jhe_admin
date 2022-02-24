/*
 * @Descripttion: 平台抽水表
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:43:35
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 16:24:13
 */

package table

// 平台游戏抽水表
type GameCommission struct {
	//用户UID
	UID int `gorm:"primaryKey" json:"uid"`
	//游戏ID
	GameId int `gorm:"primaryKey" json:"gameId,omitempty"`
	// 币种
	Symbol string `gorm:"primaryKey" json:"symbol,omitempty"`
	// 抽水
	Commission float32 `json:"commission"`
}

func (l *GameCommission) TableName() string {
	return "game_commission"
}
