package absentuser

import (
	"context"
	"errors"
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"gorm.io/gorm"
)

type AbsentUserRepository struct {
	db *gorm.DB
}

func NewAbsentUserRepository(db *gorm.DB) *AbsentUserRepository {
	return &AbsentUserRepository{db: db}
}

var _ repository.AbsentUser = (*AbsentUserRepository)(nil)

func (r *AbsentUserRepository) GetAbsentUserByUserID(ctx context.Context, userID string) (*entity.AbsentUser, error) {
	var absentUser entity.AbsentUser

	if err := r.db.WithContext(ctx).Where("user_id = ? ORDER BY created_at DESC LIMIT 1", userID).First(&absentUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &absentUser, nil
}

func (r *AbsentUserRepository) CheckInAbsentUser(ctx context.Context, absentUser *entity.AbsentUser) error {
	if err := r.db.WithContext(ctx).Create(absentUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *AbsentUserRepository) CheckOutAbsentUser(ctx context.Context, absentUser *entity.AbsentUser) error {
	if err := r.db.WithContext(ctx).Create(absentUser).Error; err != nil {
		return err
	}
	return nil
}
