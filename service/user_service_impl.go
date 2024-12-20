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
	user, err := service.UserRepository.FindByEmail(ctx, service.DB, userRequest.Email)
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
		"user_id":    user.ID,
		"user_email": user.Email,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	helper.PanicIfError(err)

	return tokenString
}

func (service *UserService) Register(ctx context.Context, userRequest web.RegisterUserRequest) web.UserResponse {
	// tx := service.DB.Begin()
	// defer helper.CommitOrRollback(tx)

	_, err := service.UserRepository.FindByEmail(ctx, service.DB, userRequest.Email)
	if err == nil {
		panic(exception.NewInvalidCredentialError("Email is already registered"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	helper.PanicIfError(err)

	user := domain.UserHouse{
		Email:    userRequest.Email,
		Password: string(hash),
		Name:     userRequest.Name,
	}
	createdUser, err := service.UserRepository.Save(ctx, service.DB, user)
	helper.PanicIfError(err)

	return web.UserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}
}

func (service *UserService) GetUserById(ctx context.Context, userId string) web.UserResponse {
	user, err := service.UserRepository.FindByUserId(ctx, service.DB, userId)
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

func (service *UserService) GetMyDetail(ctx context.Context, userId string) web.MyDetailResponse {
	myDetail, err := service.UserRepository.FindMyDetail(ctx, service.DB, userId)
	if err != nil {
		if err.Error() == "user not found" {
			panic(exception.NewDataNotFoundError("user not found"))
		}
		panic(err)
	}

	return web.MyDetailResponse{
		User: web.UserResponse{
			ID:    myDetail.User.ID,
			Email: myDetail.User.Email,
			Name:  myDetail.User.Name,
		},
		Transactions: mapTransactionsToResponse(myDetail.Transactions),
		Houses:       mapHousesToResponse(myDetail.Houses),
	}
}

func mapTransactionsToResponse(transactions []domain.UserHouseTransaction) []web.TransactionResponse {
	var response []web.TransactionResponse
	for _, txn := range transactions {
		response = append(response, web.TransactionResponse{
			ID:                txn.ID,
			UserID:            txn.UserID,
			HouseID:           txn.HouseID,
			TransactionStatus: txn.TransactionStatus,
			StartDate:         txn.StartDate,
			EndDate:           txn.EndDate,
			ExpiredAt:         txn.ExpiredAt,
		})
	}
	return response
}

func mapHousesToResponse(houses []domain.House) []web.DetailHouseResponse {
	var response []web.DetailHouseResponse
	for _, house := range houses {
		response = append(response, web.DetailHouseResponse{
			ID:            house.ID,
			Latitude:      house.Latitude,
			Longitude:     house.Longitude,
			Address:       house.Address,
			Category:      house.Category,
			UnitCount:     house.UnitCount,
			PricePerMonth: house.PricePerMonth,
		})
	}
	return response
}
