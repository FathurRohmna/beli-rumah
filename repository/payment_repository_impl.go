package repository

import (
	"beli-tanah/model/domain"
	"context"

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

func (r *PaymentRepository) UpdateWalletAndTransaction(ctx context.Context, tx *gorm.DB, orderID string, amount float64) error {
	var transaction domain.Transaction
	if err := tx.WithContext(ctx).
		Where("order_id = ? AND status = ?", orderID, "pending").
		First(&transaction).Error; err != nil {
		return err
	}

	if err := tx.WithContext(ctx).
		Model(&domain.UserHouse{}).
		Where("id = ?", transaction.UserID).
		Update("wallet_amount", gorm.Expr("wallet_amount + ?", amount)).Error; err != nil {
		return err
	}

	transaction.Status = "completed"
	if err := tx.WithContext(ctx).Save(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepository) GetUserByOrderID(ctx context.Context, tx *gorm.DB, orderID string) (*domain.UserHouse, error) {
	var transaction domain.Transaction
	if err := tx.WithContext(ctx).
		Where("order_id = ?", orderID).
		First(&transaction).Error; err != nil {
		return nil, err
	}

	var user domain.UserHouse
	if err := tx.WithContext(ctx).
		Where("id = ?", transaction.UserID).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}