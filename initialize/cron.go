/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-21 11:29:54
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 10:33:20
 */
package initialize

import (
	"JHE_admin/global"
	"JHE_admin/model"

	"github.com/robfig/cron"
	"go.uber.org/zap"
)

var Dayprofit, Weekprofit, Monthprofit float64

func NewCorn() {
	dayprofit, weekprofit, monthprofit, err := model.GetGameConnfig()
	Dayprofit = dayprofit
	Weekprofit = weekprofit
	Monthprofit = monthprofit
	if err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
	}
	c := cron.New()
	c.AddFunc("0 0 2 1,11,21 * *", CalculateVipProfit)
	c.AddFunc("0 0 2 * * *", TestDay)
	c.AddFunc("0 0 3 * * 1", TestWeek)
	c.AddFunc("0 0 4 1 * ?", TestMonth)
	c.Start()
	select {}
}

func CalculateVipProfit() {
	go model.CalculateVip()
}

func TestDay() {
	go model.CalculateGameRank_Day(Dayprofit)
}
func TestWeek() {
	go model.CalculateGameRank_Day(Weekprofit)
}
func TestMonth() {
	go model.CalculateGameRank_Day(Monthprofit)
}
