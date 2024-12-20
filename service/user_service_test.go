package service

import (
	"beli-tanah/exception"
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"beli-tanah/repository"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestLogin_Success(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{})

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	mockUser := domain.UserHouse{
		Email:    "test@examp.com",
		Password: string(hashedPassword),
		Name:     "Test User",
		ID:       "1",
	}
	userRepoMock.Mock.On("FindByEmail", mock.Anything, mock.Anything, "test@examp.com").Return(mockUser, nil)

	userService := NewUserService(userRepoMock, db)

	userRequest := web.LoginUserRequest{Email: "test@examp.com", Password: "password"}
	result := userService.Login(context.Background(), userRequest)
	assert.NotEmpty(t, result)
}

func TestLogin_UserNotFound(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{})

	userService := NewUserService(userRepoMock, db)

	userRepoMock.Mock.On("FindByEmail", mock.Anything, mock.Anything, "test@examp.com").Return(domain.UserHouse{}, errors.New("user not found"))

	userRequest := web.LoginUserRequest{Email: "test@examp.com", Password: "password"} // Use the expected email

	assert.PanicsWithValue(t, exception.NewDataNotFoundError("Username not found"), func() {
		userService.Login(context.Background(), userRequest)
	})
}

func TestLogin_IncorrectPassword(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{}) // Dummy DB

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correct_password"), 10)
	user := domain.UserHouse{
		ID:           "1",
		Name:         "Test User",
		Email:        "test@example.com",
		Password:     string(hashedPassword),
		WalletAmount: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	userRepoMock.Mock.On("FindByEmail", mock.Anything, mock.Anything, "test@example.com").Return(user, nil)

	userService := NewUserService(userRepoMock, db)

	userRequest := web.LoginUserRequest{Email: "test@example.com", Password: "wrong_password"}

	assert.PanicsWithValue(t, exception.NewInvalidCredentialError("Username or password is not valid"), func() {
		userService.Login(context.Background(), userRequest)
	})
}

func TestRegister_Success(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{})

	userService := NewUserService(userRepoMock, db)

	userRepoMock.Mock.On("FindByEmail", mock.Anything, mock.Anything, "newuser@example.com").Return(domain.UserHouse{}, errors.New("user not found"))
	userRepoMock.Mock.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("domain.UserHouse")).Return(domain.UserHouse{
		ID:    "1",
		Email: "newuser@example.com",
		Name:  "New User",
	}, nil)

	userRequest := web.RegisterUserRequest{
		Email:    "newuser@example.com",
		Password: "password",
		Name:     "New User",
	}

	result := userService.Register(context.Background(), userRequest)

	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "newuser@example.com", result.Email)
	assert.Equal(t, "New User", result.Name)
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{})

	userService := NewUserService(userRepoMock, db)

	userRepoMock.Mock.On("FindByEmail", mock.Anything, mock.Anything, "existing@example.com").Return(domain.UserHouse{
		ID:    "1",
		Email: "existing@example.com",
	}, nil)

	userRequest := web.RegisterUserRequest{
		Email:    "existing@example.com",
		Password: "password",
		Name:     "Existing User",
	}

	assert.PanicsWithValue(t, exception.NewInvalidCredentialError("Email is already registered"), func() {
		userService.Register(context.Background(), userRequest)
	})
}

func TestGetUserById_Success(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{})

	userService := NewUserService(userRepoMock, db)

	userRepoMock.Mock.On("FindByUserId", mock.Anything, mock.Anything, "1").Return(domain.UserHouse{
		ID:    "1",
		Email: "user@example.com",
		Name:  "Test User",
	}, nil)

	result := userService.GetUserById(context.Background(), "1")

	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "user@example.com", result.Email)
	assert.Equal(t, "Test User", result.Name)
}

func TestGetUserById_UserNotFound(t *testing.T) {
	userRepoMock := new(repository.IUserRepositoryMock)
	db, _ := gorm.Open(nil, &gorm.Config{})

	userService := NewUserService(userRepoMock, db)

	userRepoMock.Mock.On("FindByUserId", mock.Anything, mock.Anything, "1").Return(domain.UserHouse{}, errors.New("user not found"))

	assert.PanicsWithValue(t, exception.NewDataNotFoundError("user not found"), func() {
		userService.GetUserById(context.Background(), "1")
	})
}
