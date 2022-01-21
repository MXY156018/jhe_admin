/**
 * @Author: Dong
 * @Description:获得当前月，当前周，当前季度的初始和结束日期
 * @File:  tools
 * @Date: 2020/08/06 16:24
 */
package utils

import (
	"time"
)

/**
 * @Author Dong
 * @Description 获得当前月的初始和结束日期
 * @Date 16:29 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetMonthDay(date time.Time) (string, string) {

	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得当前周的初始和结束日期
 * @Date 16:32 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetWeekDay(date time.Time) (string, string) {
	offset := int(time.Monday - date.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - date.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description //获得当前季度的初始和结束日期
 * @Date 16:33 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetQuarterDay() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}
