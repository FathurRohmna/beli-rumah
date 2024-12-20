package repository

import (
	"beli-tanah/model/domain"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type IUserHouseTransactionRepositoryMock struct {
	mock.Mock
}

func (m *IUserHouseTransactionRepositoryMock) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction domain.UserHouseTransaction) (domain.UserHouseTransaction, error) {
	args := m.Mock.Called(ctx, tx, transaction)
	if args.Get(0) == nil {
		return domain.UserHouseTransaction{}, args.Error(1)
	}
	return args.Get(0).(domain.UserHouseTransaction), args.Error(1)
}

func (m *IUserHouseTransactionRepositoryMock) CancelTransaction(ctx context.Context, tx *gorm.DB, transactionID string) error {
	args := m.Mock.Called(ctx, tx, transactionID)
	return args.Error(0)
}

func (m *IUserHouseTransactionRepositoryMock) ConfirmTransaction(ctx context.Context, tx *gorm.DB, transactionID string) error {
	args := m.Mock.Called(ctx, tx, transactionID)
	return args.Error(0)
}

func (m *IUserHouseTransactionRepositoryMock) FindTransactionById(ctx context.Context, tx *gorm.DB, transactionID string) (domain.UserHouseTransaction, error) {
	args := m.Mock.Called(ctx, tx, transactionID)
	if args.Get(0) == nil {
		return domain.UserHouseTransaction{}, args.Error(1)
	}
	return args.Get(0).(domain.UserHouseTransaction), args.Error(1)
}
