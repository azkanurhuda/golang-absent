package handler

import (
	"github.com/azkanurhuda/golang-absent/application/usecase"
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"github.com/azkanurhuda/golang-absent/domain/service"
	"github.com/azkanurhuda/golang-absent/interfaces/middleware"
	"github.com/azkanurhuda/golang-absent/interfaces/presenter"
	"github.com/azkanurhuda/golang-absent/interfaces/response"
	"github.com/google/uuid"
	"net/http"
)

type AbsentUserHandler struct {
	absentUserUseCase usecase.AbsentUserUseCase
}

func NewAbsentUserHandler(absentUserRepository repository.AbsentUser, jwtService service.Jwt) *AbsentUserHandler {
	return &AbsentUserHandler{
		absentUserUseCase: usecase.NewAbsentUserUseCase(absentUserRepository, jwtService),
	}
}

func (h *AbsentUserHandler) CheckIn(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	payload := &entity.AbsentUser{
		ID:     uuid.NewString(),
		UserID: userID,
		Status: "Check In",
	}

	res, err := h.absentUserUseCase.CheckInAbsentUser(r.Context(), payload)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAbsentUser(res))
}

func (h *AbsentUserHandler) CheckOut(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	payload := &entity.AbsentUser{
		ID:     uuid.NewString(),
		UserID: userID,
		Status: "Check Out",
	}

	res, err := h.absentUserUseCase.CheckOutAbsentUser(r.Context(), payload)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAbsentUser(res))
}
