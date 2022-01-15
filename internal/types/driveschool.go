package types

import "time"

type DrivingSchool struct {
	Id         int       `json:"id,optional"`
	SchoolName string    `json:"school_name,optional"`
	Address    string    `json:"address,optional"`
	CreateTime time.Time `json:"create_time,optional"`
	DeleteTime time.Time `json:"delete_time,optional"`
}
