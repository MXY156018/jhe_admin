package types

import "time"

type Reward struct {
	Id         int       `json:"id,optional"`
	Uid        int       `json:"uid,optional"`
	Balance    float64   `json:"balance,optional"`
	Reward     float64   `json:"reward,optional"`
	CreateTime time.Time `json:"create_time,optional"`
	Status     int       `json:"status,optional"`
}

type RewardReq struct {
	Uid       int    `json:"uid,optional"`
	StartTime string `json:"start_time,optional"`
	EndTime   string `json:"end_time,optional"`
	Status    string `json:"status,optional"`
	PageInfo
}

type DailyReward struct {
	Num   int64   `json:"num"`
	Total float64 `json:"total"`
}
