package domain

type MidtransCallback struct {
	OrderID           string  `json:"order_id"`
	GrossAmount       float64 `json:"gross_amount,string"`
	TransactionStatus string  `json:"transaction_status"`
	SignatureKey      string  `json:"signature_key"`
}
