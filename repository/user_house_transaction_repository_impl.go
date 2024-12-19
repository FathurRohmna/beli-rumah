package repository

import (
	"beli-tanah/model/domain"
	"context"
	"errors"
	"fmt"
	"time"

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

func (r *UserHouseTransactionRepository) ConfirmTransaction(ctx context.Context, tx *gorm.DB, transactionID string) error {
	tx = tx.WithContext(ctx)

	var userTransaction domain.UserHouseTransaction
	var house domain.House
	var user domain.UserHouse

	if err := tx.First(&userTransaction, "id = ?", transactionID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find transaction: %v", err)
	}

	if err := tx.First(&house, "id = ?", userTransaction.HouseID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find house: %v", err)
	}

	if err := tx.First(&user, "id = ?", userTransaction.UserID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find user: %v", err)
	}

	if user.WalletAmount < float64(house.Size) {
		tx.Rollback()
		return fmt.Errorf("insufficient wallet balance")
	}

	user.WalletAmount -= float64(house.Size)

	userTransaction.TransactionStatus = "sold"

	houseKey := domain.HouseKey{
		TransactionID: transactionID,
		CreatedAt:     time.Now(),
	}

	if err := tx.Create(&houseKey).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create house key: %v", err)
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update user wallet: %v", err)
	}

	if err := tx.Save(&userTransaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update transaction status: %v", err)
	}

	return tx.Commit().Error
}

func (r *UserHouseTransactionRepository) FindTransactionById(ctx context.Context, tx *gorm.DB, transactionID string) (domain.UserHouseTransaction, error) {
	var transaction domain.UserHouseTransaction
	err := tx.WithContext(ctx).Where("id = ?", transactionID).First(&transaction).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.UserHouseTransaction{}, errors.New("transaction not found")
		}
		return domain.UserHouseTransaction{}, err
	}
	return transaction, nil
}
