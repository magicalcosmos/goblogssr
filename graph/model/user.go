package model

import "time"

// User User
type User struct {
	ID         int       `json:"id" db:"id"`
	Username   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Email      string    `json:"email" db:"email"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type UserWithPage struct {
	Users []User `json:"users"`
	Page  Q      `json:"page"`
}
