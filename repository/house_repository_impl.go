package repository

import (
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
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

func (r *HouseRepository) GetHouses(ctx context.Context, tx *gorm.DB, category web.HouseCategory, page, limit int) ([]domain.House, int64, error) {
	var houses []domain.House
	var totalCount int64
	tx = tx.WithContext(ctx)

	offset := (page - 1) * limit

	query := tx.Model(&domain.House{})

	if category != "" {
		query = query.Where("category = ?", category)
	}

	query = query.Offset(offset).Limit(limit)

	err := query.Find(&houses).Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	return houses, totalCount, nil
}

func (r *HouseRepository) GetHouseWithTransactions(ctx context.Context, tx *gorm.DB, houseID string) (domain.House, []domain.UserHouseTransaction, error) {
	var house domain.House
	var transactions []domain.UserHouseTransaction

	tx = tx.WithContext(ctx)

	if err := tx.Where("id = ?", houseID).First(&house).Error; err != nil {
		return domain.House{}, nil, err
	}

	if err := tx.Where("house_id = ?", houseID).
		Where("transaction_status != ?", "sold").
		Find(&transactions).Error; err != nil {
		return domain.House{}, nil, err
	}

	return house, transactions, nil
}
