package types

import "time"

type Notice struct {
	Id         int       `json:"id,optional"`
	Title      string    `json:"title,optional"`
	Content    string    `json:"content,optional"`
	CreateTime time.Time `json:"create_time,optional"`
}

type NoticePage struct {
	Notice
	PageInfo
}
