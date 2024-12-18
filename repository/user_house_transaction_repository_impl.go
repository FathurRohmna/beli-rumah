package repository

import (
	"beli-tanah/model/domain"
	"context"

	"gorm.io/gorm"
)

type UserHouseTransactionRepository struct{}

func NewUserHouseTransactionRepository() IUserHouseTransactionRepository {
	return &UserHouseTransactionRepository{}
}

func (r *UserHouseTransactionRepository) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction domain.UserHouseTransaction) (domain.UserHouseTransaction, error) {
	if err := tx.WithContext(ctx).Create(&transaction).Error; err != nil {
		return domain.UserHouseTransaction{}, err
	}

	return transaction, nil
}
