package response

import (
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"time"
)

type AccessToken struct {
	UserID    string `json:"user_id"`
	Token     string `json:"access_token"`
	ExpiresAt string `json:"expires_at"`
}

func NewAccessToken(e *entity.AccessToken) AccessToken {
	return AccessToken{
		UserID:    e.UserID,
		Token:     e.Token,
		ExpiresAt: e.ExpiresAt.Format(time.RFC3339),
	}
}
