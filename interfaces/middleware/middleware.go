package middleware

import "github.com/azkanurhuda/golang-absent/domain/service"

type Middleware struct {
	jwtService service.Jwt
}

func NewHandler(jwtService service.Jwt) *Middleware {
	return &Middleware{
		jwtService: jwtService,
	}
}
