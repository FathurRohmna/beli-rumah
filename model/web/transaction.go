package web

import "time"

type TopUpUserWalletGeneratePaymentResponse struct {
	PaymentUrl string `json:"payment_url"`
}

type UserHouseTransactionResponse struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	HouseID string `json:"house_id"`
	Status  string `json:"status"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`

	ExpiredAt time.Time `json:"expired_at"`
}

type UserHouseTransactionPreviewResponse struct {
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type TopUpUserWalletGeneratePaymentRequest struct {
	Amount float64 `json:"amount"`
}
