package web

import "time"

type UserResponse struct {
	ID           string  `json:"id"`
	Email        string  `json:"email"`
	Name         string  `json:"name"`
	WalletAmount float64 `json:"wallet_amount"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"example_password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"example_password"`
}

type MyDetailResponse struct {
	User         UserResponse          `json:"user"`
	Transactions []TransactionResponse `json:"transactions"`
	Houses       []DetailHouseResponse `json:"houses"`
}

type TransactionResponse struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	HouseID           string    `json:"house_id"`
	TransactionStatus string    `json:"transaction_status"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	ExpiredAt         time.Time `json:"expired_at"`
}

type DetailHouseResponse struct {
	ID            string        `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Latitude      float64       `json:"latitude"`
	Longitude     float64       `json:"longitude"`
	Address       string        `json:"address"`
	Category      HouseCategory `json:"category"`
	UnitCount     int           `json:"unit_count"`
	PricePerMonth float64       `json:"price_per_month"`
}
