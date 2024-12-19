package repository

import (
	"beli-tanah/model/domain"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HouseKeyRepository struct{}

func NewHouseKeyRepository() IHouseKeyRepository {
	return &HouseKeyRepository{}
}

func (r *HouseKeyRepository) CountActiveHouseKeys(ctx context.Context, tx *gorm.DB, houseID string, startDate time.Time, endDate time.Time) (int64, error) {
	var count int64

	if err := tx.WithContext(ctx).
		Model(&domain.HouseKey{}).
		Joins("JOIN user_house_transactions ON house_keys.transaction_id = user_house_transactions.id").
		Where("user_house_transactions.house_id = ? AND user_house_transactions.start_date <= ? AND (user_house_transactions.end_date IS NULL OR user_house_transactions.end_date >= ?)", houseID, startDate, endDate).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("error counting active house keys: %v", err)
	}

	return count, nil
}
