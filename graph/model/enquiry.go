package model

import "time"

// User User
type Enquiry struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Phone    string    `json:"phone" db:"phone"`
	Type     string    `json:"type" db:"type"`
	Message  string    `json:"message" db:"message"`
	CreateAt time.Time `json:"createAt" db:"create_time"`
}

type EnquiryWithPage struct {
	Enquiry []Enquiry `json:"enquiries"`
	Page    Q         `json:"page"`
}
