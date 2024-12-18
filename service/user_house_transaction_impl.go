package service

import (
	"beli-tanah/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserHouseTransactionService struct {
	UserHouseTransactionRepository repository.IUserHouseTransactionRepository
	DB                             *gorm.DB
}

func NewUserHouseTransactionService(userHouseTransactionRepository repository.IUserHouseTransactionRepository, DB *gorm.DB) IUserHouseTransactionService {
	return &UserHouseTransactionService{
		UserHouseTransactionRepository: userHouseTransactionRepository,
		DB:                             DB,
	}
}

func (service *UserHouseTransactionService) CancelTransaction(ctx context.Context, userID, transactionID string) error {
	tx := service.DB.Begin()

	if err := service.UserHouseTransactionRepository.CancelTransaction(ctx, tx, transactionID); err != nil {
		return fmt.Errorf("failed to cancel transaction: %v", err)
	}
	return nil
}

func (service *UserHouseTransactionService) ConfirmTransaction(ctx context.Context, userID, transactionID string) error {
	tx := service.DB.Begin()

	if err := service.UserHouseTransactionRepository.ConfirmTransaction(ctx, tx, transactionID); err != nil {
		return fmt.Errorf("failed to confirm transaction: %v", err)
	}
	return nil
}
