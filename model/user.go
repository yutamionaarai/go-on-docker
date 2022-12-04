package model

import "time"

// A User is ...
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Todos     []Todo    `json:"todos"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
