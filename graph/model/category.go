package model

import "time"

// User User
type Category struct {
	ID         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Pid        string    `json:"pid" db:"pid"`
	Sort       string    `json:"sort" db:"sort"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
