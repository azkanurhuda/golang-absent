package repository

import (
	"context"
	"github.com/azkanurhuda/golang-absent/domain/entity"
)

type AbsentUser interface {
	GetAbsentUserByUserID(ctx context.Context, userID string) (*entity.AbsentUser, error)
	CheckInAbsentUser(ctx context.Context, absentUser *entity.AbsentUser) error
	CheckOutAbsentUser(ctx context.Context, absentUser *entity.AbsentUser) error
}
