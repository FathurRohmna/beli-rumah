package service

import (
	"beli-tanah/helper"
	"beli-tanah/model/domain"
	"beli-tanah/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HouseService struct {
	HouseRepository                repository.IHouseRepository
	UserHouseTransactionRepository repository.IUserHouseTransactionRepository
	DB                             *gorm.DB
}

func NewHouseService(houseRepository repository.IHouseRepository, userHouseTransactionRepository repository.IUserHouseTransactionRepository, DB *gorm.DB) IHouseService {
	return &HouseService{
		HouseRepository:                houseRepository,
		UserHouseTransactionRepository: userHouseTransactionRepository,
		DB:                             DB,
	}
}

func (service *HouseService) BuyHouseTransaction(ctx context.Context, userID, houseID string) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	house, err := service.HouseRepository.FindHouseByID(ctx, tx, houseID)
	if err != nil {
		return fmt.Errorf("house not found: %v", err)
	}

	pendingCount, err := service.HouseRepository.CountPendingTransactions(ctx, tx, houseID)
	if err != nil {
		return fmt.Errorf("error checking pending transactions: %v", err)
	}

	if pendingCount >= int64(house.UnitCount) {
		return fmt.Errorf("no available slots, please wait until another transaction completes")
	}

	expiryTime := time.Now().Add(1 * time.Minute)
	transaction := domain.UserHouseTransaction{
		UserID:            userID,
		HouseID:           houseID,
		TransactionStatus: "pending",
		ExpiredAt:         expiryTime,
	}

	_, err = service.UserHouseTransactionRepository.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %v", err)
	}

	return nil
}

func (service *HouseService) CancelTransaction(ctx context.Context, userID, transactionID string) error {
	tx := service.DB.Begin()

	if err := service.UserHouseTransactionRepository.CancelTransaction(ctx, tx, transactionID); err != nil {
		return fmt.Errorf("failed to cancel transaction: %v", err)
	}
	return nil
}

func (service *HouseService) ConfirmTransaction(ctx context.Context, userID, transactionID string) error {
	tx := service.DB.Begin()

	if err := service.HouseRepository.ConfirmTransaction(ctx, tx, transactionID); err != nil {
		return fmt.Errorf("failed to confirm transaction: %v", err)
	}
	return nil
}
