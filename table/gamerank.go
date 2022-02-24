/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 16:35:36
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 03:42:18
 */

package table

// type GameRank struct {
// 	Date   string  `gorm:"primaryKey" json:"date"`
// 	UID    int     `gorm:"primaryKey" json:"uid"`
// 	GameID int     `gorm:"primaryKey" json:"gameId"`
// 	Credit float32 `json:"credit"`
// }

type GameRankToday struct {
	Date   string  `gorm:"primaryKey" json:"date"`
	UID    int     `gorm:"primaryKey" json:"uid"`
	GameID int     `gorm:"primaryKey" json:"gameId"`
	Credit float32 `json:"credit"`
}

func (l *GameRankToday) TableName() string {
	return "game_rank_today"
}

type GameRankWeek struct {
	Date   string  `gorm:"primaryKey" json:"date"`
	UID    int     `gorm:"primaryKey" json:"uid"`
	GameID int     `gorm:"primaryKey" json:"gameId"`
	Credit float32 `json:"credit"`
}

func (l *GameRankWeek) TableName() string {
	return "game_rank_week"
}

type GameRankMonth struct {
	Date   string  `gorm:"primaryKey" json:"date"`
	UID    int     `gorm:"primaryKey" json:"uid"`
	GameID int     `gorm:"primaryKey" json:"gameId"`
	Credit float32 `json:"credit"`
}

func (l *GameRankMonth) TableName() string {
	return "game_rank_month"
}
