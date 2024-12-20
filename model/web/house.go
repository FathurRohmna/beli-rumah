package web

import (
	"time"
)

type HouseCategory string

const (
	Apartment          HouseCategory = "apartment"
	Villa              HouseCategory = "villa"
	House              HouseCategory = "house"
	ResidentialComplex HouseCategory = "residentialComplex"
)

type BuyHouseResponse struct {
	TransactionToken string `json:"transaction_token"`
}

type HouseResponse struct {
	ID            string        `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Latitude      float64       `json:"latitude"`
	Longitude     float64       `json:"longitude"`
	Address       string        `json:"address"`
	Category      HouseCategory `json:"category"`
	UnitCount     int           `json:"unit_count"`
	PricePerMonth float64       `json:"price_per_month"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type HouseDetailResponse struct {
	ID                    string                                `json:"id"`
	Latitude              float64                               `json:"latitude"`
	Longitude             float64                               `json:"longitude"`
	Address               string                                `json:"address"`
	Category              HouseCategory                         `json:"category"`
	UnitCount             int                                   `json:"unit_count"`
	PricePerMonth         float64                               `json:"price_per_month"`
	CreatedAt             time.Time                             `json:"created_at"`
	UpdatedAt             time.Time                             `json:"updated_at"`
	UserHouseTransactions []UserHouseTransactionPreviewResponse `json:"user_house_transactions"`
}

type BuyHouseTransactionRequest struct {
	HouseID   string `json:"house_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
