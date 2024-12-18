package service

import "context"

type IHouseService interface {
	BuyHouseTransaction(ctx context.Context, userID, houseID string) error
}
