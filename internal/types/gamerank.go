/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-01-19 10:29:19
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 14:13:39
 */
package types

type GameRankConfig struct {
	Id     int     `json:"id"`
	Num    float64 `json:"num"`
	Remark string  `json:"remark"`
	Symbol string  `json:"symbol"`
}

type GameConfig struct {
	Day       float64 `json:"day"`
	Week      float64 `json:"week"`
	Month     float64 `json:"month"`
	SumProfit float64 `json:"sum_profit"`
}

type GameRankSearch struct {
	Uid  int    `json:"uid,optional"`
	Time string `json:"time,optional"`
}

type GameRankList struct {
	Start_time string  `json:"start_time,omitempty"`
	End_time   string  `json:"end_time,omitempty"`
	Uid        int     `json:"uid,omitempty"`
	GameId     int     `json:"game_id,omitempty"`
	Credit     float64 `json:"credit"`
	Rank       int     `json:"rank,omitempty"`
	Profit     float64 `json:"profit"`
}

type RankList struct {
	DayList   []*GameRankList `json:"day_list"`
	WeekList  []*GameRankList `json:"week_list"`
	MonthList []*GameRankList `json:"month_list"`
}
