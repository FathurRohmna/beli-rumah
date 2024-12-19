package repository

import (
	"beli-tanah/model/domain"
	"context"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.UserHouse) (domain.UserHouse, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.UserHouse, error)
	FindByUserId(ctx context.Context, tx *gorm.DB, userId string) (domain.UserHouse, error)
}

