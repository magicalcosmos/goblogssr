package model

import (
	"database/sql"
	"time"
)

// User User
type Category struct {
	ID         string         `json:"id" db:"id"`
	Name       string         `json:"name" db:"name"`
	ParentId   sql.NullString `json:"parentId" db:"parent_id"`
	Sort       string         `json:"sort" db:"sort"`
	Status     int            `json:"status" db:"status"`
	CreateTime time.Time      `json:"createAt" db:"create_time"`
	UpdateTime time.Time      `json:"updateAt" db:"update_time"`
}
