package repository

import (
	"beli-tanah/model/domain"
	"context"

	"gorm.io/gorm"
)

type IPaymentRepository interface {
	TopUpUserWalletTransaction(ctx context.Context, tx *gorm.DB, transaction domain.Transaction) (domain.Transaction, error)
	UpdateWalletAndTransaction(ctx context.Context, tx *gorm.DB, transactionId string, amount float64) error
}
