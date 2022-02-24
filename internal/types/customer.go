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
	Issue      string    `json:"issue,omitempty"`       //對戰回合編號
	Mode       string    `json:"mode,omitempty"`        //對戰模式
	Symbol     string    `json:"symbol,omitempty"`      //幣種
	Ticket     int       `json:"ticket,omitempty"`      //門票
	Rebate     int       `json:"rebate,omitempty"`      //抽水
	Uid        int       `json:"uid,omitempty"`         //用戶uid
	Camp       string    `json:"camp,omitempty"`        //陣營
	Death      int       `json:"death,omitempty"`       //陣亡次數
	Result     string    `json:"result,omitempty"`      //結果
	PreBalance float64   `json:"pre_balance,omitempty"` //戰前餘額
	CurBalance float64   `json:"cur_balance,omitempty"` //對戰後餘額
	Win        int       `json:"win,omitempty"`         //應分
	CreateTime time.Time `json:"create_time"`           //时间
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
