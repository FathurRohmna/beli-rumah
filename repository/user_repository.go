package repository

import (
	"beli-tanah/model/domain"
	"context"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error)
}

