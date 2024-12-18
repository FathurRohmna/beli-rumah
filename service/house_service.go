package service

import "context"

type IHouseService interface {
	BuyHouseTransaction(ctx context.Context, userID, houseID string) error
	CancelTransaction(ctx context.Context, userID, transactionID string) error
	ConfirmTransaction(ctx context.Context, userID, transactionID string) error
}
