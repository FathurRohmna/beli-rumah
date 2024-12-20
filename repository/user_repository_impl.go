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

func (r *UserRepository) FindMyDetail(ctx context.Context, tx *gorm.DB, userID string) (domain.MyDetail, error) {
	var user domain.UserHouse
	var transactions []domain.UserHouseTransaction
	var houseIDs []string
	var houses []domain.House

	err := tx.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.MyDetail{}, errors.New("user not found")
		}
		return domain.MyDetail{}, err
	}

	err = tx.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&transactions).Error
	if err != nil {
		return domain.MyDetail{}, err
	}

	for _, transaction := range transactions {
		houseIDs = append(houseIDs, transaction.HouseID)
	}

	err = tx.WithContext(ctx).
		Where("id IN ?", houseIDs).
		Find(&houses).Error
	if err != nil {
		return domain.MyDetail{}, err
	}

	myDetail := domain.MyDetail{
		User:         user,
		Transactions: transactions,
		Houses:       houses,
	}

	return myDetail, nil
}
