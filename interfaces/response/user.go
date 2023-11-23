package response

import "github.com/azkanurhuda/golang-absent/domain/entity"

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(e *entity.User) User {
	return User{
		UserID:   e.ID,
		Username: e.Username,
		Email:    e.Email,
	}
}
