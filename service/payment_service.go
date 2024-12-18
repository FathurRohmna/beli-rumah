package service

import (
	"context"
	"beli-tanah/model/web"
)

type IPaymentService interface {
	TopUpUserWalletGeneratePayment(ctx context.Context, userID string, amount float64) web.TopUpUserWalletGeneratePaymentResponse
}
