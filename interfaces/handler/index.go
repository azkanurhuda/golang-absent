package handler

import (
	"github.com/azkanurhuda/golang-absent/interfaces/presenter"
	"gorm.io/gorm"
	"net/http"
)

type IndexHandler struct {
	db *gorm.DB
}

func NewIndexHandler(db *gorm.DB) *IndexHandler {
	return &IndexHandler{db: db}
}

func (h *IndexHandler) Index(w http.ResponseWriter, _ *http.Request) {
	presenter.OK(w)
}

func (h *IndexHandler) Healthy(w http.ResponseWriter, _ *http.Request) {
	presenter.OK(w)
}
