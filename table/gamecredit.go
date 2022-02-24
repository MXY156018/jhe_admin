/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:44:10
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 03:42:33
 */

package table

// type GameCredit struct {
// 	Date   string  `gorm:"primaryKey" json:"date"`
// 	UID    int     `gorm:"primaryKey" json:"uid"`
// 	Credit float32 `json:"credit"`
// }

type GameCreditToday struct {
	Date   string  `gorm:"primaryKey" json:"date"`
	UID    int     `gorm:"primaryKey" json:"uid"`
	Credit float32 `json:"credit"`
}

func (l *GameCreditToday) TableName() string {
	return "game_credit_today"
}

type GameCreditWeek struct {
	Date   string  `gorm:"primaryKey" json:"date"`
	UID    int     `gorm:"primaryKey" json:"uid"`
	Credit float32 `json:"credit"`
}

func (l *GameCreditWeek) TableName() string {
	return "game_credit_week"
}

type GameCreditMonth struct {
	Date   string  `gorm:"primaryKey" json:"date"`
	UID    int     `gorm:"primaryKey" json:"uid"`
	Credit float32 `json:"credit"`
}

func (l *GameCreditMonth) TableName() string {
	return "game_credit_month"
}
