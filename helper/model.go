package helper

import (
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
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
