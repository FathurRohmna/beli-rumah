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
