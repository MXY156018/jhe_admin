/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 16:19:35
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 15:42:28
 */
/**
 * @Author: Dong
 * @Description:获得当前月，当前周，当前季度的初始和结束日期
 * @File:  tools
 * @Date: 2020/08/06 16:24
 */
package utils

import (
	"fmt"
	"time"
)

// 获取本日时间范围
//
// date 日期
//
// onlyDate 是否值返回日期根式，而不返回时间
//
// onlyDate = true 返回本日日期
//
// onlyDate = false 返回 本日时间,本日结束时间
func GetToday(date time.Time, onlyDate bool) (string, string) {
	currentYear, currentMonth, currentDay := date.Date()
	currentLocation := date.Location()
	today := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
	if onlyDate {
		return today.Format("2006-01-02"), ""
	}
	return today.Format("2006-01-02") + " 00:00:00", today.Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得当前月的初始和结束日期
 * @Date 16:29 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetMonthDay(date time.Time, onlyDate bool) (string, string) {

	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	if onlyDate {
		return time.Unix(f, 0).Format("2006-01-02"), time.Unix(l, 0).Format("2006-01-02")
	}
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得当前周的初始和结束日期
 * @Date 16:32 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetWeekDay(date time.Time, onlyDate bool) (string, string) {
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
	if onlyDate {
		return time.Unix(f, 0).Format("2006-01-02"), time.Unix(l, 0).Format("2006-01-02")
	}
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
func GetTenDay(today time.Time) (time.Time, time.Time) {
	currentYear, currentMonth, currentDay := today.Date()
	currentLocation := today.Location()
	today = time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
	ad, _ := time.ParseDuration("240h")
	day_end := today.Add(ad)
	return today, day_end
}

//获取设置用户总充值的时间节点
func GetSumDate() string {
	day := time.Now().Day()

	if day >= 21 {
		day = 21
	} else if day >= 11 {
		day = 11
	} else if day >= 1 {
		day = 1
	}
	year := time.Now().Year()
	month := time.Now().Month()

	dateStr := fmt.Sprintf("%d-%d-%d", year, int(month), day)
	return dateStr
}

func GetPreDate(now time.Time) string {
	var day, month, year int
	day = now.Day()
	fmt.Println("day", day)
	if day >= 21 {

		month = int(now.Month())
		year = now.Year()
		day = 11
		fmt.Println("时间大于21号")
	} else if day >= 11 {

		month = int(now.Month())
		year = now.Year()
		day = 1
		fmt.Println("时间大于11号")
	} else if day >= 1 {

		nowmonth := now.Month()
		nowyear := now.Year()
		if nowmonth == 1 {
			year = nowyear - 1
			month = 12
			day = 21
		} else {
			year = nowyear
			month = int(nowmonth) - 1
			day = 21
		}
		fmt.Println("时间大于1号")
	}
	dateStr := fmt.Sprintf("%d-%d-%d", year, int(month), day)
	return dateStr
}

func GetPreWeek(now time.Time) string {
	ad, _ := time.ParseDuration("-168h")
	now = now.Add(ad)

	start, _ := GetWeekDay(now, true)
	return start
}
func GetPreDay(now time.Time) string {
	ad, _ := time.ParseDuration("-24h")
	now = now.Add(ad)

	start, _ := GetToday(now, true)
	return start
}
func GetPreMonth(now time.Time) string {
	year := now.Year()
	var month int
	if now.Month() == 1 {
		month = 12
		year -= 1
	} else {
		month = int(now.Month()) - 1
	}
	return fmt.Sprintf("%d-%d-%d", year, int(month), 1)
}
