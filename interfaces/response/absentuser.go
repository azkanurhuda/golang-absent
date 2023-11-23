package response

import "github.com/azkanurhuda/golang-absent/domain/entity"

type AbsentUser struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	IPAddress string  `json:"ip_address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}

func NewAbsentUser(e *entity.AbsentUser) AbsentUser {
	return AbsentUser{
		ID:        e.ID,
		UserID:    e.UserID,
		IPAddress: e.IPAddress,
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Status:    e.Status,
	}
}
