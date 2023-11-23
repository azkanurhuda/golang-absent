package router

import (
	"github.com/azkanurhuda/golang-absent/interfaces/handler"
	"github.com/azkanurhuda/golang-absent/interfaces/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	r.HandleFunc("/", h.Index).Methods("GET")
	r.HandleFunc("/healthy", h.Healthy).Methods("GET")
	r.HandleFunc("/signup", h.Signup).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(h.JWT())
	v1.HandleFunc("/me", h.Me).Methods("GET")

	v1.HandleFunc("/checkin", h.CheckIn).Methods("POST")
	v1.HandleFunc("/checkout", h.CheckOut).Methods("POST")

	return r
}
