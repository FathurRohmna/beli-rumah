package service

import (
	"beli-tanah/exception"
	"beli-tanah/helper"
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"beli-tanah/repository"
	"context"
	"os"

	"github.com/golang-jwt/jwt"
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

	return web.UserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}
}

func (service *UserService) GetUserById(ctx context.Context, userId string) web.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		if err.Error() == "user not found" {
			panic(exception.NewDataNotFoundError("user not found"))
		}

		panic(err)
	}

	return web.UserResponse{
		ID:    userId,
		Email: user.Email,
		Name:  user.Name,
	}
}
