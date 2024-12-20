package repository

import (
	"beli-tanah/model/domain"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type IUserRepositoryMock struct {
	Mock mock.Mock
}

func (m *IUserRepositoryMock) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.UserHouse, error) {
	args := m.Mock.Called(ctx, tx, email)
	if args.Get(0) == nil {
		return domain.UserHouse{}, args.Error(1)
	}
	return args.Get(0).(domain.UserHouse), args.Error(1)
}
func (m *IUserRepositoryMock) FindByUserId(ctx context.Context, tx *gorm.DB, userId string) (domain.UserHouse, error) {
	args := m.Mock.Called(ctx, tx, userId)
	return args.Get(0).(domain.UserHouse), args.Error(1)
}

func (m *IUserRepositoryMock) Save(ctx context.Context, tx *gorm.DB, user domain.UserHouse) (domain.UserHouse, error) {
	args := m.Mock.Called(ctx, tx, user)
	return args.Get(0).(domain.UserHouse), args.Error(1)
}

func (m *IUserRepositoryMock) FindMyDetail(ctx context.Context, tx *gorm.DB, userID string) (domain.MyDetail, error) {
	args := m.Mock.Called(ctx, tx, userID)
	if args.Get(0) == nil {
		return domain.MyDetail{}, args.Error(1)
	}
	return args.Get(0).(domain.MyDetail), args.Error(1)
}
