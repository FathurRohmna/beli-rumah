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

func (r *UserRepository) Save(ctx context.Context, tx *gorm.DB, user domain.UserHouse) (domain.UserHouse, error) {
	err := tx.WithContext(ctx).Save(&user).Error
	if err != nil {
		return domain.UserHouse{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.UserHouse, error) {
	var user domain.UserHouse
	err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.UserHouse{}, errors.New("user not found")
		}
		return domain.UserHouse{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userId string) (domain.UserHouse, error) {
	var user domain.UserHouse
	err := tx.WithContext(ctx).Where("id = ?", userId).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.UserHouse{}, errors.New("user not found")
		}
		return domain.UserHouse{}, err
	}
	return user, nil
}
