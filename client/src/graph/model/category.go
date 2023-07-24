package model

// Category category
type Category struct {
	ID     int64    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Pid    int    `json:"pid" db:"pid"`
	Sort   int    `json:"sort" db:"sort"`
	Status int    `json:"status" db:"status"`
}

// NewCategory new category for gql
type NewCategory struct {
	ID string `json:"id"`
	Name string `json:"name"`
}
