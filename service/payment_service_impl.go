package service

import (
	"beli-tanah/config"
	"beli-tanah/helper"
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"beli-tanah/repository"
	"context"
	"crypto/sha512"
	"fmt"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type PaymentService struct {
	PaymentRepository repository.IPaymentRepository
	DB                *gorm.DB
}

func NewPaymentService(paymentRepository repository.IPaymentRepository, DB *gorm.DB) IPaymentService {
	return &PaymentService{
		PaymentRepository: paymentRepository,
		DB:                DB,
	}
}

func (service *PaymentService) TopUpUserWalletGeneratePayment(ctx context.Context, userID string, amount float64) web.TopUpUserWalletGeneratePaymentResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	orderID := "TOPUP-" + time.Now().Format("20060102150405")

	transaction := domain.Transaction{
		UserID:  userID,
		OrderID: orderID,
		Amount:  amount,
		Status:  "pending",
	}
	transactionRes, err := service.PaymentRepository.TopUpUserWalletTransaction(ctx, tx, transaction)
	helper.PanicIfError(err)

	client := config.SetupMidtrans()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionRes.OrderID,
			GrossAmt: int64(transactionRes.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: "frohman@students.hacktiv8.ac.id",
		},
	}

	snapResp, _ := client.CreateTransaction(req)

	return web.TopUpUserWalletGeneratePaymentResponse{
		PaymentUrl: snapResp.RedirectURL,
	}
}

func (service *PaymentService) UpdateWalletAndTransaction(ctx context.Context, transactionID string, amount float64) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	err := service.PaymentRepository.UpdateWalletAndTransaction(ctx, tx, transactionID, amount)
	helper.PanicIfError(err)

	return nil
}

func (service *PaymentService) VerifyMidtransSignature(callback domain.MidtransCallback) bool {
	serverKey := config.SetupMidtrans().ServerKey
	data := callback.OrderID + fmt.Sprintf("%.0f", callback.GrossAmount) + serverKey

	expectedSignature := fmt.Sprintf("%x", sha512.Sum512([]byte(data)))
	return expectedSignature == callback.SignatureKey
}
