package repository

import (
	"beli-tanah/model/domain"
	"context"

	"gorm.io/gorm"
)

type IUserHouseTransactionRepository interface {
	CreateTransaction(ctx context.Context, tx *gorm.DB, transaction domain.UserHouseTransaction) (domain.UserHouseTransaction, error)
	CancelTransaction(ctx context.Context, tx *gorm.DB, transactionID string) error
}
