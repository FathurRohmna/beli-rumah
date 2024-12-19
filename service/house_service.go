package service

import (
	"beli-tanah/model/web"
	"context"
	"time"
)

type IHouseService interface {
	CheckPaymentAvailability(ctx context.Context, houseID string, startDate time.Time, endDate time.Time) error
	CheckHouseAvailability(ctx context.Context, houseID string, startDate time.Time, endDate time.Time) error
	BuyHouseTransaction(ctx context.Context, userID, houseID string, startDate time.Time, endDate time.Time) (web.BuyHouseResponse, error)
}
