package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type IHouseKeyRepository interface {
	CountActiveHouseKeys(ctx context.Context, tx *gorm.DB, houseID string, startDate time.Time, endDate time.Time) (int64, error)
}
