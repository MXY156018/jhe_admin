/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-11 15:16:41
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-23 17:14:00
 */
package types

import "time"

type HallUser struct {
	Account  string `json:"account"`
	PassWord string `json:"passWord"`
}
type User struct {
	Uid           int64     `json:"uid"`
	Head          int       `json:"head"`
	Parent        int       `json:"parent"`
	Account       string    `json:"account"`
	Address       string    `json:"address"`
	RegisterTime  time.Time `json:"registerTime"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	LastLoginIP   string    `json:"lastLoginIp"`
	IsBot         int       `json:"isBot"`
	Type          string    `json:"type"`
}
type Wallet struct {
	Uid            int     `json:"uid,omitempty"`
	Symbol         string  `json:"symbol,optional"`
	Balance        float64 `json:"balance,optional"`
	GameFrozen     float64 `json:"game_frozen,omitempty,optional"`
	WithdeawFrozen float64 `json:"withdeaw_frozen,omitempty,optional"`
}

type WalletAddress struct {
	Wallet  []Wallet `json:"wallet"`
	Address string   `json:"address"`
}
type LoginResp struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	AccessToken  string      `json:"token"`
	AccessExpire int64       `json:"accessExpire"`
	RefreshAfter int64       `json:"refreshAfter"`
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
type UserWithdrawl struct {
	Id         int       `json:"id,omitempty"`
	Uid        int       `json:"uid,omitempty"`         //用户id
	Status     string    `json:"status,omitempty"`      //状态 0 提币请求 1 提币中 2 成功 3 失败 4 已经退回
	ReqDate    time.Time `json:"req_date,omitempty"`    //请求日期
	FinishDate time.Time `json:"finish_date,omitempty"` //完成日期
	Symbol     string    `json:"symbol,omitempty"`      //币种
	Amount     float64   `json:"amount"`                //实际到账额度
	TxHash     string    `json:"tx_hash,omitempty"`     //交易HASH
	RawAmount  float64   `json:"raw_amount,omitempty"`  //原始请求额度
	Fee        float64   `json:"fee,omitempty"`         //手续费
	To         string    `json:"to,omitempty"`          //转到哪个地址
	PreBalance float64   `json:"pre_balance"`           //提币前余额
	NowBalance float64   `json:"now_balance"`           //提币后余额
}

type RechargeSum struct {
	Uid         int            `json:"uid"`
	Parent      int            `json:"parent"`
	Type        string         `json:"type"`
	Status      int            `json:"status"`
	Children    []*RechargeSum `json:"children" gorm:"foreignKey:uid;references:uid;"`
	SumRecharge float64        `json:"sum_recharge"`
}
type FeedBack struct {
	Id         int       `json:"id,optional"`
	Uid        int       `json:"uid,optional"`
	Message    string    `json:"message,optional"`
	Handle     string    `json:"handle,optional"`
	CreateTime time.Time `json:"create_time,optional"`
}
type GameConfig struct {
	Day    float64 `json:"day"`
	Week   float64 `json:"week"`
	Month  float64 `json:"month"`
	Symbol string  `json:"symbol"`
}
type GameRankList struct {
	Rank   int     `json:"rank"`
	Uid    int     `json:"uid"`
	Credit float32 `json:"credit"`
	Num    float32 `json:"num,omitempty"`
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
	Uid        int       `json:"uid,optional"`
	Type       int       `json:"type,optional"`
	Num        float64   `json:"num"`
	PreBalance float64   `json:"pre_balance,optional,omitempty"`
	Balance    float64   `json:"balance,optional,omitempty"`
	CreateTime time.Time `json:"create_time,optional"`
	DrawTime   time.Time `json:"draw_time,omitempty"`
	IsDraw     int       `json:"is_draw,omitempty"`
	Symbol     string    `json:"Symbol,omitempty"`
	SumNum     float64   `json:"sum_num"`
}

type CustomerPage struct {
	Uid      int `json:"uid,optional"`
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

type SysConfig struct {
	RechargeFee  string `json:"recharge_fee,omitempty"`
	WithdrawlFee string `json:"withdrawl_fee,omitempty"`
	GameCost     string `json:"game_cost,omitempty"`
	GameFee      string `json:"game_fee,omitempty"`
}
type GameRankConfig struct {
	Id     int     `json:"id"`
	Num    float64 `json:"num"`
	Remark string  `json:"remark"`
}

type VipConfig struct {
	BuyVipSymbol string  `json:"buy_vip_symbol,omitempty"`
	BuyVipCost   float32 `json:"buy_vip_cost,omitempty"`
}

type UserBuyVip struct {
	Id         int       `json:"id,omitempty"`
	Uid        int       `json:"uid,omitempty"`
	Num        float32   `json:"num,omitempty"`
	Symbol     string    `json:"symbol,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}
