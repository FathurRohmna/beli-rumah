package service

import (
	"beli-tanah/model/web"
	"context"
)

type IUserHouseTransactionService interface {
	CancelTransaction(ctx context.Context, userID, transactionID string) error
	ConfirmTransaction(ctx context.Context, userID, transactionID string) error
	FindTransactionById(ctx context.Context, transactionID string) (web.UserHouseTransactionResponse, error)
}
