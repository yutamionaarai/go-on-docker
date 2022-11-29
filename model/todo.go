package model

import "time"

type Todo struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	Priority       int64     `json:"priority"`
	ExpirationDate time.Time `json:"expiration_date"`
	UserID         int64     `json:"user_id"`
	CreatedAT      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
