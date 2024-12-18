package service

import (
	"beli-tanah/exception"
	"beli-tanah/helper"
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"beli-tanah/repository"
	"context"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository repository.IUserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.IUserRepository, DB *gorm.DB) *UserService {
	return &UserService{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *UserService) Login(ctx context.Context, userRequest web.LoginUserRequest) string {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, userRequest.Email)
	if err != nil {
		if err.Error() == "user not found" {
			panic(exception.NewDataNotFoundError("Username not found"))
		}

		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		panic(exception.NewInvalidCredentialError("Username or password is not valid"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	helper.PanicIfError(err)

	return tokenString
}

func (service *UserService) Register(ctx context.Context, userRequest web.RegisterUserRequest) web.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := service.UserRepository.FindByEmail(ctx, tx, userRequest.Email)
	if err == nil {
		panic(exception.NewInvalidCredentialError("Email is already registered"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	helper.PanicIfError(err)

	user := domain.User{
		Email:    userRequest.Email,
		Password: string(hash),
		Name:     userRequest.Name,
	}
	createdUser, err := service.UserRepository.Save(ctx, tx, user)
	helper.PanicIfError(err)

	userID, err := uuid.Parse(createdUser.ID)
	if err != nil {
		log.Fatalf("Invalid user ID: %v", err)
	}

	return web.UserResponse{
		ID:    userID,
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}
}
