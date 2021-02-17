package model

// User User
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUser new user for gql
type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
