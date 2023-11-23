package usecase

import (
	"context"
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"github.com/azkanurhuda/golang-absent/domain/service"
	"github.com/azkanurhuda/golang-absent/pkg/crypt"
	log "github.com/sirupsen/logrus"
)

type UserUseCase interface {
	SignUp(ctx context.Context, user *entity.User) (*entity.AccessToken, error)
	Login(ctx context.Context, email, password string) (*entity.AccessToken, error)
	Me(ctx context.Context, userID string) (*entity.User, error)
}

type userUseCase struct {
	userRepository repository.User
	jwtService     service.Jwt
}

func NewUserUseCase(userRepository repository.User, jwtService service.Jwt) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (u *userUseCase) SignUp(ctx context.Context, user *entity.User) (*entity.AccessToken, error) {
	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	token, err := u.jwtService.Sign(user)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return token, nil
}

func (u *userUseCase) Login(ctx context.Context, email, password string) (*entity.AccessToken, error) {
	user, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}
	if user == nil {
		return nil, newNotFoundError()
	}

	if err = crypt.ComparePassword(user.Password, password); err != nil {
		log.Error(err)
		return nil, newErrorUnauthorized()
	}

	token, err := u.jwtService.Sign(user)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return token, nil
}

func (u *userUseCase) Me(ctx context.Context, userID string) (*entity.User, error) {
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}
	if user == nil {
		return nil, newNotFoundError()
	}

	return user, nil
}
