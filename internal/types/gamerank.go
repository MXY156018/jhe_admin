package types

import "time"

type GameRankConfig struct {
	Id     int     `json:"id"`
	Num    float64 `json:"num"`
	Remark string  `json:"remark"`
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
	Rank      int       `json:"rank"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Uid       int       `json:"uid"`
	SumScore  int       `json:"sum_score"`
	Profit    float64   `json:"profit"`
}

type RankList struct {
	DayList   []*GameRankList `json:"day_list"`
	WeekList  []*GameRankList `json:"week_list"`
	MonthList []*GameRankList `json:"month_list"`
}
