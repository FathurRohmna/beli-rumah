package service

import (
	"beli-tanah/exception"
	"beli-tanah/model/web"
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
	if err := service.UserHouseTransactionRepository.CancelTransaction(ctx, service.DB, transactionID); err != nil {
		return fmt.Errorf("failed to cancel transaction: %v", err)
	}
	return nil
}

func (service *UserHouseTransactionService) ConfirmTransaction(ctx context.Context, userID, transactionID string) error {
	if err := service.UserHouseTransactionRepository.ConfirmTransaction(ctx, service.DB, transactionID); err != nil {
		return fmt.Errorf("failed to confirm transaction: %v", err)
	}
	return nil
}

func (service *UserHouseTransactionService) FindTransactionById(ctx context.Context, transactionID string) (web.UserHouseTransactionResponse, error) {
	transaction, err := service.UserHouseTransactionRepository.FindTransactionById(ctx, service.DB, transactionID)
	if err != nil {
		if err.Error() == "transaction not found" {
			panic(exception.NewDataNotFoundError("transaction not found"))
		}

		panic(err)
	}

	return web.UserHouseTransactionResponse{
		ID:        transaction.ID,
		UserID:    transaction.UserID,
		Status:    transaction.TransactionStatus,
		HouseID:   transaction.HouseID,
		StartDate: transaction.StartDate,
		EndDate:   transaction.EndDate,
		ExpiredAt: transaction.ExpiredAt.Local(),
	}, nil
}
