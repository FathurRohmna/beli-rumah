package repository

import (
	"beli-tanah/model/domain"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HouseRepository struct{}

func NewHouseRepository() IHouseRepository {
	return &HouseRepository{}
}

func (r *HouseRepository) FindHouseByID(ctx context.Context, tx *gorm.DB, houseID string) (domain.House, error) {
	var house domain.House

	if err := tx.WithContext(ctx).First(&house, "id = ?", houseID).Error; err != nil {
		return house, fmt.Errorf("house not found: %v", err)
	}

	return house, nil
}

func (r *HouseRepository) CountPendingTransactions(ctx context.Context, tx *gorm.DB, houseID string) (int64, error) {
	var count int64
	now := time.Now()

	if err := tx.WithContext(ctx).Model(&domain.UserHouseTransaction{}).
		Where("house_id = ? AND transaction_status = 'pending' AND expired_at > ?", houseID, now).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("error counting pending transactions: %v", err)
	}

	return count, nil
}

func (r *HouseRepository) ConfirmTransaction(ctx context.Context, tx *gorm.DB, transactionID string) error {
	tx = tx.WithContext(ctx).Begin()

	var userTransaction domain.UserHouseTransaction
	var house domain.House
	var user domain.User

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

	if house.UnitCount <= 0 {
		tx.Rollback()
		return fmt.Errorf("not enough units available")
	}

	house.UnitCount -= 1
	user.WalletAmount -= float64(house.Size)

	userTransaction.TransactionStatus = "sold"

	if err := tx.Save(&house).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update house: %v", err)
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
