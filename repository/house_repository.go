package repository

import (
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"context"
	"time"

	"gorm.io/gorm"
)

type IHouseRepository interface {
	FindHouseByID(ctx context.Context, tx *gorm.DB, houseID string) (domain.House, error)
	CountPendingTransactions(ctx context.Context, tx *gorm.DB, houseID string, startDate, endDate time.Time) (int64, error)
	GetHouses(ctx context.Context, tx *gorm.DB, category web.HouseCategory, latitude, longitude float64, page, limit int) ([]domain.House, int64, error)
}
