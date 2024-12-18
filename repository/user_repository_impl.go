package repository

import (
	"beli-tanah/model/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	err := tx.WithContext(ctx).Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}
