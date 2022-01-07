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
}
