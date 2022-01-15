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
