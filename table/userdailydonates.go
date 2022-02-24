/*
 * @Descripttion: user_daily_donate_beans 表
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 01:35:13
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 09:50:52
 */

package table

import "time"

type UserDailyDonate struct {
	// 用户UID
	UID int `gorm:"primaryKey json:"uid"`
	// 上次赠送日期
	Date time.Time `json:"date"`
	// 额度
	Amount float32 `json:"amount"`
}

func (l *UserDailyDonate) TableName() string {
	return "user_daily_donates"
}
