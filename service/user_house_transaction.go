package service

import "context"

type IUserHouseTransactionService interface {
	CancelTransaction(ctx context.Context, userID, transactionID string) error
	ConfirmTransaction(ctx context.Context, userID, transactionID string) error
}
