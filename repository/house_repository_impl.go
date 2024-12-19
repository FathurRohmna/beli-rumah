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

func (r *HouseRepository) CountPendingTransactions(ctx context.Context, tx *gorm.DB, houseID string, startDate, endDate time.Time) (int64, error) {
	var count int64

	err := tx.WithContext(ctx).
		Model(&domain.UserHouseTransaction{}).
		Where("house_id = ?", houseID).
		Where("transaction_status = ?", "pending").
		Where("expired_at > CURRENT_TIMESTAMP").
		Where(
			"("+
				"(end_date IS NULL AND ? >= start_date) OR "+
				"(end_date IS NOT NULL AND "+
				"("+
				"(start_date <= ? AND end_date >= ?) OR "+
				"(start_date >= ? AND start_date <= ?)"+
				")"+
				")"+
				")", endDate, endDate, startDate, startDate, endDate,
		).
		Count(&count).Error

	if err != nil {
		return 0, fmt.Errorf("error counting pending transactions: %v", err)
	}

	return count, nil
}
