package helper

import (
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"fmt"
	"log"
	"time"
)

func MapDomainToBuyHouseResponse(houses []domain.House) []web.HouseResponse {
	var response []web.HouseResponse
	for _, house := range houses {
		response = append(response, web.HouseResponse{
			ID:            house.ID,
			Latitude:      house.Latitude,
			Longitude:     house.Longitude,
			Address:       house.Address,
			Category:      house.Category,
			UnitCount:     house.UnitCount,
			PricePerMonth: house.PricePerMonth,
			CreatedAt:     house.CreatedAt,
			UpdatedAt:     house.UpdatedAt,
		})
	}
	return response
}

func ParseDate(dateStr string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return time.Time{}, fmt.Errorf("invalid date format, expected yyyy-mm-dd")
	}
	return parsedDate, nil
}
