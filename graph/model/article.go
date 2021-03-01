package model

import "time"

// Article article
type Article struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	Title      string    `json:"title" db:"title"`
	CategoryID int64     `json:"category_id" db:"category_id"`
	Tag        string    `json:"tag" db:"tag"`
	Remark     string    `json:"remark" db:"remark"`
	Desc       string    `json:"desc" db:"desc"`
	PV         int64     `json:"pv" db:"pv"`
	Created    time.Time `json:"created" db:"created"`
	Updated    time.Time `json:"updated" db:"updated"`
	Status     int       `json:"status" db:"status"`
	Review     int       `json:"review" db:"review"`
}
