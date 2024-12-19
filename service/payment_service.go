package service

import (
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"context"
)

type IPaymentService interface {
	TopUpUserWalletGeneratePayment(ctx context.Context, userID string, amount float64) web.TopUpUserWalletGeneratePaymentResponse
	UpdateWalletAndTransaction(ctx context.Context, transactionID string, amount float64) error
	VerifyMidtransSignature(callback domain.MidtransCallback) bool
}
