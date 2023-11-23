package form

import (
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/pkg/crypt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SignUp struct {
	Username string `form:"username" validate:"required,max=40"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

func (f *SignUp) Validate() error {
	validate := validator.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}

func (f *SignUp) Entity() (*entity.User, error) {
	hash, err := crypt.GenerateBCryptoHash(f.Password)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       uuid.NewString(),
		Username: f.Username,
		Email:    f.Email,
		Password: hash,
	}, nil
}

type Login struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

func (f *Login) Validate() error {
	validate := validator.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}
