package repository

import (
	"beli-tanah/model/domain"
	"context"
	"fmt"

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

func (r *UserHouseTransactionRepository) CancelTransaction(ctx context.Context, tx *gorm.DB, transactionID string) error {
	if err := tx.WithContext(ctx).Model(&domain.UserHouseTransaction{}).
		Where("id = ? AND transaction_status = 'pending'", transactionID).
		Update("transaction_status", "cancelled").Error; err != nil {
		return fmt.Errorf("failed to cancel transaction: %v", err)
	}

	return nil
}
