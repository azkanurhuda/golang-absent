package user

import (
	"context"
	"errors"
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

var _ repository.User = (*UserRepository)(nil)

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
