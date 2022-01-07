package types

import "time"

type Equipments struct {
	Id            int       `json:"id,optional"`
	Name          string    `json:"name,optional"`
	EId           string    `json:"e_id,optional"`
	Status        int       `json:"status,optional"`
	CreateTime    time.Time `json:"create_time,optional"`
	DeleteTime    time.Time `json:"delete_time,optional"`
	IsUpdate      int       `json:"is_update,optional"`
	DrivingSchool int       `json:"driving_school,optional"`
	SchoolName    string    `json:"school_name,optional"`
}
type EquipmentList struct {
	Equipments
	PageInfo
}

type ChangeEquipments struct {
	Id            int       `json:"id,optional"`
	Name          string    `json:"name,optional"`
	EId           string    `json:"e_id,optional"`
	Status        int       `json:"status,optional"`
	CreateTime    time.Time `json:"create_time,optional"`
	DeleteTime    time.Time `json:"delete_time,optional"`
	IsUpdate      int       `json:"is_update,optional"`
	DrivingSchool int       `json:"driving_school,optional"`
}
