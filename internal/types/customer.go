/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 10:56:58
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 11:20:40
 */
package types

import "time"

type User struct {
	Uid           int       `json:"uid,omitempty,optional"`
	Head          int       `json:"head,omitempty,optional"`
	Parent        int       `json:"parent,omitempty,optional"`
	Account       string    `json:"account,omitempty,optional"`
	Password      string    `json:"-"`
	Address       string    `json:"address,omitempty,optional"`
	RegisterTime  time.Time `json:"register_time,omitempty,optional"`
	LastLoginTime time.Time `json:"last_login_time,omitempty,optional"`
	LastLoginIP   string    `json:"last_login_ip,omitempty,optional"`
	IsBot         int       `json:"is_bot,omitempty,optional"`
	Type          string    `json:"type,omitempty,optional"`
	Status        int       `json:"status,omitempty,optional"`
}
type UserReq struct {
	Uid    int `json:"uid,omitempty,optional"`
	Status int `json:"status,omitempty,optional"`
}
type UserRecharge1 struct {
	SumRecharge    float64 `json:"sum_recharge,optional"`
	SubSumRecharge float64 `json:"sub_sum_recharge,optional"`
}
type UserRecharge struct {
	Id         int       `json:"id,omitempty"`
	Uid        int       `json:"uid,omitempty"`
	Date       time.Time `json:"date,omitempty"`
	Symbol     string    `json:"symbol,omitempty"`
	Amount     float64   `json:"amount,omitempty"`
	TxHash     string    `json:"tx_hash,omitempty"`
	RawAmount  float64   `json:"raw_amount,omitempty"`
	Fee        float64   `json:"fee,omitempty"`
	PreBalance float64   `json:"pre_balance"`
	NowBalance float64   `json:"now_balance"`
}
type UserDetail struct {
	User
	UserRecharge1
}
type Wallet struct {
	Uid            int     `json:"uid,omitempty,optional"`
	Symbol         string  `json:"symbol,omitempty,optional"`
	Balance        float64 `json:"balance,optional"`
	GameFrozen     float64 `json:"game_frozen,optional"`
	WithdrawFrozen float64 `json:"withdraw_frozen,optional"`
}
type UserList struct {
	PageInfo
	List []User `json:"list"`
}
type CustimerSearch struct {
	Uid       int    `json:"uid,omitempty,optional"`
	Type      string `json:"type,omitempty,optional"`
	Status    string `json:"status,omitempty,optional"`
	StartTime string `json:"startTime,optional"`
	EndTime   string `json:"endTime,optional"`
	PageInfo
}

type GameRecord struct {
	Id         int       `json:"id,omitempty"`
	Uid        int       `json:"uid,omitempty"`
	Type       int       `json:"type,omitempty"`
	Num        float64   `json:"num,omitempty"`
	PreBalance float64   `json:"pre_balance,omitempty"`
	Balance    float64   `json:"balance,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}

type GameRecordG1 struct {
	Issue      string    `json:"issue,omitempty"`       //??????????????????
	Mode       string    `json:"mode,omitempty"`        //????????????
	Symbol     string    `json:"symbol,omitempty"`      //??????
	Ticket     int       `json:"ticket,omitempty"`      //??????
	Rebate     int       `json:"rebate,omitempty"`      //??????
	Uid        int       `json:"uid,omitempty"`         //??????uid
	Camp       string    `json:"camp,omitempty"`        //??????
	Death      int       `json:"death,omitempty"`       //????????????
	Result     string    `json:"result,omitempty"`      //??????
	PreBalance float64   `json:"pre_balance,omitempty"` //????????????
	CurBalance float64   `json:"cur_balance,omitempty"` //???????????????
	Win        int       `json:"win,omitempty"`         //??????
	CreateTime time.Time `json:"create_time"`           //??????
}

type GameRecordList struct {
	GameRecordG1
	PageInfo
}

type CustomerOperator struct {
	Uid        int       `json:"uid,optional"`
	Type       int       `json:"type,optional"`
	Num        float64   `json:"num,optional"`
	PreBalance float64   `json:"pre_balance,optional"`
	Balance    float64   `json:"balance,optional"`
	CreateTime time.Time `json:"create_time,optional"`
	DrawTime   time.Time `json:"draw_time,optional"`
	IsDraw     int       `json:"is_draw,optional"`
	Symbol     string    `json:"symbol,optional"`
	SumNum     float64   `json:"sum_num,optional"`
}

type OperateRecord struct {
	CustomerOperator
	PageInfo
}

type BillReq struct {
	Uid       int    `json:"uid,optional"`
	Type      int    `json:"type,optional"`
	StartTime string `json:"startTime,optional"`
	EndTime   string `json:"endTime,optional"`
	PageInfo
}

type GameDailyRep struct {
	Recharge float64 `json:"recharge"`
	Reward   float64 `json:"reward"`
	Platform float64 `json:"platform"`
}

type HomeData struct {
	UserNum        int64   `json:"user_num"`
	PlatformProfit float64 `json:"platform_profit"`
	RechargeNum    float64 `json:"recharge_num"`
	RewardNum      float64 `json:"reward_num"`
	VipbalanceNum  float64 `json:"vipbalance_num"`
	GameablanceNum float64 `json:"gameablance_num"`
}
