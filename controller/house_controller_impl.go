package controller

import (
	"beli-tanah/model/web"
	"beli-tanah/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type HouseController struct {
	HouseService service.IHouseService
	EmailService service.IEmailService
}

func NewHouseController(houseService service.IHouseService, emailService service.IEmailService) IHouseController {
	return &HouseController{HouseService: houseService, EmailService: emailService}
}

func (controller *HouseController) BuyHouseTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	startDate, err := time.Parse("02-01-2006", "09-12-2024")
	if err != nil {
		log.Fatalf("Failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("02-01-2006", "10-12-2024")
	if err != nil {
		log.Fatalf("Failed to parse end date: %v", err)
	}

	token, err := controller.HouseService.BuyHouseTransaction(ctx,
		"0e4f279d-03ea-46dc-a07a-92d057e1e470",
		"792d35f3-92cb-4c45-bad0-2042ab02c4aa",
		startDate,
		endDate,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	controller.EmailService.SendEmail(ctx, "frohman@students.hacktiv8.ac.id", "Testing email here", token.TransactionToken)

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction cancelled"})
}

func (controller *HouseController) GetHouses(c echo.Context) error {
	category := c.QueryParam("category")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	latitude, _ := strconv.ParseFloat(c.QueryParam("latitude"), 64)
	longitude, _ := strconv.ParseFloat(c.QueryParam("longitude"), 64)

	var houseCategory web.HouseCategory
	switch category {
	case "apartment":
		houseCategory = web.Apartment
	case "villa":
		houseCategory = web.Villa
	case "house":
		houseCategory = web.House
	case "residentialComplex":
		houseCategory = web.ResidentialComplex
	default:
		return c.JSON(http.StatusBadRequest, "Invalid category")
	}

	houses, totalCount, err := controller.HouseService.GetHouses(c.Request().Context(), houseCategory, latitude, longitude, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"houses":     houses,
		"totalCount": totalCount,
	})
}
