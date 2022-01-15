package types

import "time"

type Customers struct {
	Id         int       `json:"id,omitempty,optional"`
	Sid        int       `json:"sid,omitempty,optional"`
	Address    string    `json:"address,omitempty,optional"`
	Type       string    `json:"type,omitempty,optional"`
	CreateTime time.Time `json:"create_time,omitempty,optional"`
	Status     int       `json:"status,omitempty,optional"`
}
type Wallet struct {
	Wid        int     `json:"wid,omitempty,optional"`
	Uid        int     `json:"uid,omitempty,optional"`
	Currencyid int     `json:"currencyid,omitempty,optional"`
	Name       string  `json:"name,omitempty,optional"`
	Balance    float64 `json:"balance,omitempty,optional"`
	Lock       float64 `json:"lock,omitempty,optional"`
}
type CustomerList struct {
	PageInfo
	Customers
	Wallet
	SumSubordinateRecharge float64 `json:"sum_subordinate_recharge"`
}

type CustimerSearch struct {
	Id        int    `json:"id,omitempty,optional"`
	Type      string `json:"type,omitempty,optional"`
	Status    string `json:"status,omitempty,optional"`
	StartTime string `json:"startTime,optional"`
	EndTime   string `json:"endTime,optional"`
	PageInfo
}

type GameRecord struct {
	Id         int       `json:"id,omitempty,optional"`
	Uid        int       `json:"uid,omitempty,optional"`
	Gameid     int       `json:"gameid,omitempty,optional"`
	Type       int       `json:"type,optional"`
	PreBalance float64   `json:"pre_balance,optional"`
	NowBalance float64   `json:"now_balance,optional"`
	Win        float64   `json:"win,optional"`
	Status     int       `json:"status,optional"`
	Commission float64   `json:"commission,optional"`
	StartTime  time.Time `json:"start_time,optional"`
	EndTime    time.Time `json:"end_time,optional"`
}

type GameRecordList struct {
	GameRecord
	PageInfo
}

type CustomerOperator struct {
	Id         int       `json:"id,optional"`
	Uid        int       `json:"uid,optional"`
	Type       int       `json:"type,optional"`
	Num        float64   `json:"num,optional"`
	Balance    float64   `json:"balance,optional"`
	CreateTime time.Time `json:"create_time,optional"`
}

type OperateRecord struct {
	CustomerOperator
	PageInfo
}

type BillReq struct {
	Uid       int    `json:"uid,optional"`
	Type      int    `json:"type,optional"`
	StartTime string `json:"start_time,optional"`
	EndTime   string `json:"end_time,optional"`
	PageInfo
}

type GameDailyRep struct {
	Recharge float64 `json:"recharge"`
	Reward   float64 `json:"reward"`
	Platform float64 `json:"platform"`
}
