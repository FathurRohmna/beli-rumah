package service

import (
	"beli-tanah/model/web"
	"context"
)

type IHouseService interface {
	CheckHouseAvailability(ctx context.Context, houseID string) error 
	BuyHouseTransaction(ctx context.Context, userID, houseID string) (web.BuyHouseResponse, error)
}
