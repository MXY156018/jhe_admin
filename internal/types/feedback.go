package types

import "time"

type FeedBack struct {
	Id          int        `json:"id,optional"`
	Uid         int        `json:"uid,optional"`
	Message     string     `json:"message,optional"`
	CreateTime  time.Time  `json:"create_time,optional,"`
	ResolveTime *time.Time `json:"resolve_time,optional"`
	DeleteTime  *time.Time `json:"delete_time,optional"`
	Phone       string     `json:"phone,optional"`
	Email       string     `json:"email,optional"`
	Status      string     `json:"status,optional"`
	Handle      string     `json:"handle,optional"`
	IsRead      int        `json:"is_read,optional"`
	UIsRead     int        `json:"u_is_read,optional"`
}
type SearchFeedBack struct {
	Uid       int    `json:"uid,optional"`
	Message   string `json:"message,optional"`
	StartTime string `json:"start_time,optional"`
	EndTime   string `json:"end_time,optional"`
	Phone     string `json:"phone,optional"`
	Email     string `json:"email,optional"`
	Status    string `json:"status,optional"`
}
type FeedBackList struct {
	FeedBack
	PageInfo
}

type FeedBackReq struct {
	SearchFeedBack
	PageInfo
}
