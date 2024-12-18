package repository

import (
	"context"
	"beli-tanah/model/domain"

	"gorm.io/gorm"
)

type PaymentRepository struct{}

func NewPaymentRepository() IPaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) TopUpUserWalletTransaction(ctx context.Context, tx *gorm.DB, transaction domain.Transaction) (domain.Transaction, error) {
	if err := tx.WithContext(ctx).Create(&transaction).Error; err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}
