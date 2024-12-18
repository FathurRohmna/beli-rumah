package web

type TopUpUserWalletGeneratePaymentResponse struct {
	PaymentUrl string `json:"payment_url"`
}

type UserHouseTransactionResponse struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	HouseID string `json:"house_id"`
	Status  string `json:"status"`
}
