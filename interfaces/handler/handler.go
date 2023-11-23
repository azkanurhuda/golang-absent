package handler

import (
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"github.com/azkanurhuda/golang-absent/domain/service"
	"github.com/azkanurhuda/golang-absent/interfaces/middleware"
	"gorm.io/gorm"
)

type Handler struct {
	*IndexHandler
	*UserHandler
	*AuthenticationHandler
	*middleware.Middleware
	*AbsentUserHandler
}

func NewHandler(db *gorm.DB,
	userRepository repository.User,
	jwtService service.Jwt,
	absentUserRepository repository.AbsentUser,
) *Handler {
	indexHandler := NewIndexHandler(db)
	userHandler := NewUserHandler(userRepository, jwtService)
	authenticationHandler := NewAuthenticationHandler(userRepository, jwtService)
	absentUserHandler := NewAbsentUserHandler(absentUserRepository, jwtService)

	return &Handler{
		IndexHandler:          indexHandler,
		UserHandler:           userHandler,
		AuthenticationHandler: authenticationHandler,
		Middleware:            middleware.NewHandler(jwtService),
		AbsentUserHandler:     absentUserHandler,
	}
}
