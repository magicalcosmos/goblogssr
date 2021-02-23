package model

import "time"

// User User
type User struct {
	ID       string    `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`
	Created  time.Time `json:"created" db:"created"`
	Status   int       `json:"status" db:"status"`
}

// NewUser new user for gql
type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
