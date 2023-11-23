package service

import "github.com/azkanurhuda/golang-absent/domain/entity"

type Jwt interface {
	Sign(user *entity.User) (*entity.AccessToken, error)
	Verify(signedToken string) (string, error)
}
