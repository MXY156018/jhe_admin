/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 12:30:52
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 11:34:55
 */
package types

import "time"

type UserWithdrawl struct {
	Id         int       `json:"id,optional,omitempty"`
	Uid        int       `json:"uid,optional,omitempty"` //用户id
	Status     int       `json:"status,optional"`        //状态 0 提币请求 1 提币中 2 成功 3 失败 4 已经退回
	ReqDate    time.Time `json:"req_date,omitempty"`     //请求日期
	FinishDate time.Time `json:"finish_date,omitempty"`  //完成日期
	Symbol     string    `json:"symbol"`                 //币种
	Amount     float64   `json:"amount,omitempty"`       //实际到账额度
	TxHash     string    `json:"tx_hash"`                //交易HASH
	RawAmount  float64   `json:"raw_amount"`             //原始请求额度
	Fee        float64   `json:"fee"`                    //手续费
	PreBalance float64   `json:"pre_balance"`            //提币前余额
	NowBalance float64   `json:"now_balance"`            //提币后余额
}

type RewardReq struct {
	Uid       int    `json:"uid,optional"`
	StartTime string `json:"startTime,optional"`
	EndTime   string `json:"endTime,optional"`
	Status    string `json:"status,optional"`
	PageInfo
}

type DailyReward struct {
	Num   int64   `json:"num"`
	Total float64 `json:"total"`
}
