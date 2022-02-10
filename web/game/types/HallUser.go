package types

import (
	"JHE_admin/internal/types"
	"time"
)

type Customer struct {
	Id         int       `json:"id,optional"`
	Sid        int       `json:"sid,optional"`
	Address    string    `json:"address,optional"`
	Type       string    `json:"type,optional"`
	CreateTime time.Time `json:"create_time,optional"`
	Status     string    `json:"status,optional"`
}

type Wallet struct {
	Wid        int     `json:"wid,omitempty"`
	Uid        int     `json:"uid,omitempty"`
	Currencyid int     `json:"currencyid,omitempty"`
	Name       string  `json:"name,omitempty"`
	Balance    float64 `json:"balance,omitempty"`
	Lock       float64 `json:"lock,omitempty"`
}

type CustomerOperator struct {
	Id         int       `json:"id,optional"`
	Uid        int       `json:"uid,optional"`
	Type       int       `json:"type,optional"`
	Num        float64   `json:"num,optional"`
	Balance    float64   `json:"balance,optional"`
	CreateTime time.Time `json:"create_time,optional"`
}

type Config struct {
	RankConfig   []types.GameRankConfig `json:"rank_config,omitempty"`
	ProfitConfig types.GameConfig       `json:"profit_config,omitempty"`
}
