package service

import (
	"beli-tanah/model/domain"
	"beli-tanah/repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCancelTransaction(t *testing.T) {
	mockRepo := new(repository.IUserHouseTransactionRepositoryMock)
	mockDB := &gorm.DB{}
	transactionID := "txn123"

	mockRepo.On("CancelTransaction", mock.Anything, mock.Anything, transactionID).Return(nil)

	svc := NewUserHouseTransactionService(mockRepo, mockDB)
	err := svc.CancelTransaction(context.Background(), "user1", transactionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCancelTransaction_Error(t *testing.T) {
	mockRepo := new(repository.IUserHouseTransactionRepositoryMock)
	mockDB := &gorm.DB{}
	transactionID := "txn123"

	mockRepo.On("CancelTransaction", mock.Anything, mock.Anything, transactionID).Return(errors.New("cancel failed"))

	svc := NewUserHouseTransactionService(mockRepo, mockDB)
	err := svc.CancelTransaction(context.Background(), "user1", transactionID)

	assert.Error(t, err)
	assert.Equal(t, "failed to cancel transaction: cancel failed", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestConfirmTransaction(t *testing.T) {
	mockRepo := new(repository.IUserHouseTransactionRepositoryMock)
	mockDB := &gorm.DB{}
	transactionID := "txn123"

	mockRepo.On("ConfirmTransaction", mock.Anything, mock.Anything, transactionID).Return(nil)

	svc := NewUserHouseTransactionService(mockRepo, mockDB)
	err := svc.ConfirmTransaction(context.Background(), "user1", transactionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestConfirmTransaction_Error(t *testing.T) {
	mockRepo := new(repository.IUserHouseTransactionRepositoryMock)
	mockDB := &gorm.DB{}
	transactionID := "txn123"

	mockRepo.On("ConfirmTransaction", mock.Anything, mock.Anything, transactionID).Return(errors.New("confirm failed"))

	svc := NewUserHouseTransactionService(mockRepo, mockDB)
	err := svc.ConfirmTransaction(context.Background(), "user1", transactionID)

	assert.Error(t, err)
	assert.Equal(t, "failed to confirm transaction: confirm failed", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestFindTransactionById(t *testing.T) {
	mockRepo := new(repository.IUserHouseTransactionRepositoryMock)
	mockDB := &gorm.DB{}
	transactionID := "txn123"

	expectedTransaction := domain.UserHouseTransaction{
		ID:                transactionID,
		UserID:            "user1",
		TransactionStatus: "confirmed",
		HouseID:           "house1",
	}

	mockRepo.On("FindTransactionById", mock.Anything, mock.Anything, transactionID).Return(expectedTransaction, nil)

	svc := NewUserHouseTransactionService(mockRepo, mockDB)
	resp, err := svc.FindTransactionById(context.Background(), transactionID)

	assert.NoError(t, err)
	assert.Equal(t, transactionID, resp.ID)
	assert.Equal(t, "confirmed", resp.Status)
	mockRepo.AssertExpectations(t)
}

func TestFindTransactionById_Error(t *testing.T) {
	mockRepo := new(repository.IUserHouseTransactionRepositoryMock)
	mockDB := &gorm.DB{}
	transactionID := "txn123"

	mockRepo.On("FindTransactionById", mock.Anything, mock.Anything, transactionID).Return(domain.UserHouseTransaction{}, errors.New("transaction not found"))

	svc := NewUserHouseTransactionService(mockRepo, mockDB)
	assert.Panics(t, func() {
		_, _ = svc.FindTransactionById(context.Background(), transactionID)
	})
	mockRepo.AssertExpectations(t)
}
