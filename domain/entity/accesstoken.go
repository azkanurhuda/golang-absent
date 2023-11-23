package entity

import "time"

type AccessToken struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}
