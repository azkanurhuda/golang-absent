package handler

import (
	"github.com/azkanurhuda/golang-absent/application/usecase"
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"github.com/azkanurhuda/golang-absent/domain/service"
	"github.com/azkanurhuda/golang-absent/interfaces/form"
	"github.com/azkanurhuda/golang-absent/interfaces/presenter"
	"github.com/azkanurhuda/golang-absent/interfaces/response"
	"net/http"
)

type AuthenticationHandler struct {
	userUseCase usecase.UserUseCase
}

func NewAuthenticationHandler(userRepository repository.User, jwtService service.Jwt) *AuthenticationHandler {
	return &AuthenticationHandler{
		userUseCase: usecase.NewUserUseCase(userRepository, jwtService),
	}
}

func (h *AuthenticationHandler) Signup(w http.ResponseWriter, r *http.Request) {
	f := form.SignUp{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	user, err := f.Entity()
	if err != nil {
		presenter.NewBadRequest(w)
		return
	}

	token, err := h.userUseCase.SignUp(r.Context(), user)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}

func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	f := form.Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	token, err := h.userUseCase.Login(r.Context(), f.Email, f.Password)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}
