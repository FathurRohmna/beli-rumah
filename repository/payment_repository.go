package repository

import (
	"context"
	"beli-tanah/model/domain"

	"gorm.io/gorm"
)

type IPaymentRepository interface {
	TopUpUserWalletTransaction(ctx context.Context, tx *gorm.DB, transaction domain.Transaction) (domain.Transaction, error)
}
