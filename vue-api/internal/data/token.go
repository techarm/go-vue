package data

import "time"

type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Toekn     string    `json:"toekn"`
	TokenHash []byte    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
