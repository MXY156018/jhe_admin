package types

import "time"

type RechargeSum struct {
	Id       int            `json:"id"`
	Sid      int            `json:"sid"`
	Type     string         `json:"type"`
	Status   int            `json:"status"`
	Children []*RechargeSum `json:"children" gorm:"foreignKey:id;references:id;"`

	SumRecharge float64   `json:"sum_recharge"`
	CreateTime  time.Time `json:"create_time"`
}
type FeedBack struct {
	Id         int       `json:"id,optional"`
	Uid        int       `json:"uid,optional"`
	Message    string    `json:"message,optional"`
	Handle     string    `json:"handle,optional"`
	CreateTime time.Time `json:"create_time,optional"`
}
type GameConfig struct {
	Day   float64 `json:"day"`
	Week  float64 `json:"week"`
	Month float64 `json:"month"`
}
type GameRankList struct {
	Rank     int `json:"rank"`
	Uid      int `json:"uid"`
	SumScore int `json:"sum_score"`
}
type Rank struct {
	DayRank   []GameRankList `json:"day_rank"`
	WeekRank  []GameRankList `json:"week_rank"`
	MonthRank []GameRankList `json:"month_rank"`
}

type RankList struct {
	RankList Rank       `json:"rank_list"`
	Config   GameConfig `json:"config"`
}
type CustomerOperator struct {
	Id         int       `json:"id,optional"`
	Uid        int       `json:"uid,optional"`
	Type       int       `json:"type,optional"`
	Num        float64   `json:"num,optional"`
	PreBalance float64   `json:"pre_balance,optional"`
	Balance    float64   `json:"balance,optional"`
	CreateTime time.Time `json:"create_time,optional"`
}

type CustomerPage struct {
	Uid      int `json:"uid"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type GameRank struct {
	Uid        int       `json:"uid,optional"`
	Num        float64   `json:"num,optional"`
	SumProfit  float64   `json:"sum_profit"`
	CreateTime time.Time `json:"create_time,optional"`
}
type BillList struct {
	Recharge  PageResult `json:"recharge"`
	Reward    PageResult `json:"reward"`
	GameRank  PageResult `json:"game_rank"`
	VipProfit PageResult `json:"vip_profit"`
}
