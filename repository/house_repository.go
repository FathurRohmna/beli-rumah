package repository

import (
	"beli-tanah/model/domain"
	"context"

	"gorm.io/gorm"
)

type IHouseRepository interface {
	FindHouseByID(ctx context.Context, tx *gorm.DB, houseID string) (domain.House, error)
	CountPendingTransactions(ctx context.Context, tx *gorm.DB, houseID string) (int64, error)
}
