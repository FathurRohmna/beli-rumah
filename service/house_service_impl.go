package service

import (
	"beli-tanah/helper"
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"beli-tanah/repository"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type HouseService struct {
	HouseRepository                repository.IHouseRepository
	UserHouseTransactionRepository repository.IUserHouseTransactionRepository
	HouseKeyRepository             repository.IHouseKeyRepository
	DB                             *gorm.DB
}

func NewHouseService(houseRepository repository.IHouseRepository, userHouseTransactionRepository repository.IUserHouseTransactionRepository, houseKeyRepository repository.IHouseKeyRepository, DB *gorm.DB) IHouseService {
	return &HouseService{
		HouseRepository:                houseRepository,
		UserHouseTransactionRepository: userHouseTransactionRepository,
		HouseKeyRepository:             houseKeyRepository,
		DB:                             DB,
	}
}

func (service *HouseService) CheckPaymentAvailability(ctx context.Context, houseID string, startDate time.Time, endDate time.Time) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if err := service.CheckHouseAvailability(ctx, houseID, startDate, endDate); err != nil {
		return fmt.Errorf("house is not available: %v", err)
	}

	house, err := service.HouseRepository.FindHouseByID(ctx, tx, houseID)
	if err != nil {
		return fmt.Errorf("house not found: %v", err)
	}

	activeKeys, err := service.HouseKeyRepository.CountActiveHouseKeys(ctx, tx, houseID, startDate, endDate)
	if err != nil {
		return fmt.Errorf("error checking pending transactions: %v", err)
	}

	availableSlot := int64(house.UnitCount) - activeKeys

	pendingCount, err := service.HouseRepository.CountPendingTransactions(ctx, tx, houseID, startDate, endDate)
	if err != nil {
		return fmt.Errorf("error checking pending transactions: %v", err)
	}

	if pendingCount >= availableSlot {
		return fmt.Errorf("no available slots, please wait until another transaction completes")
	}

	return nil
}

func (service *HouseService) CheckHouseAvailability(ctx context.Context, houseID string, startDate time.Time, endDate time.Time) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	house, err := service.HouseRepository.FindHouseByID(ctx, tx, houseID)
	if err != nil {
		return fmt.Errorf("house not found: %v", err)
	}

	activeKeys, err := service.HouseKeyRepository.CountActiveHouseKeys(ctx, tx, houseID, startDate, endDate)
	if err != nil {
		return fmt.Errorf("error checking pending transactions: %v", err)
	}

	if activeKeys >= int64(house.UnitCount) {
		return fmt.Errorf("no available slots, please wait until another dates")
	}

	return nil
}

func (service *HouseService) BuyHouseTransaction(ctx context.Context, userID, houseID string, startDate time.Time, endDate time.Time) (web.BuyHouseResponse, error) {
	if err := service.CheckHouseAvailability(ctx, houseID, startDate, endDate); err != nil {
		return web.BuyHouseResponse{}, err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	expiryTime := time.Now().Add(5 * time.Minute).Local().UTC()
	transaction := domain.UserHouseTransaction{
		UserID:            userID,
		HouseID:           houseID,
		StartDate:         startDate,
		EndDate:           endDate,
		TransactionStatus: "pending",
		ExpiredAt:         expiryTime,
	}

	transactionResp, err := service.UserHouseTransactionRepository.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return web.BuyHouseResponse{}, fmt.Errorf("failed to create transaction: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":        transaction.UserID,
		"transaction_id": transactionResp.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_AUTH_EMAIL_URL")))
	helper.PanicIfError(err)

	return web.BuyHouseResponse{
		TransactionToken: tokenString,
	}, nil
}
