package controller

import (
	"beli-tanah/model/web"
	"beli-tanah/service"
	"fmt"
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
	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing user ID", "message": "UNAUTHORIZED"})
	}

	userEmail, ok := c.Get("user_email").(string)
	if !ok || userEmail == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing user ID", "message": "UNAUTHORIZED"})
	}

	ctx := c.Request().Context()
	var request web.BuyHouseTransactionRequest

	if err := c.Bind(&request); err != nil {
		fmt.Print(request)

		return c.JSON(http.StatusBadRequest, map[string]string{"info": "Invalid request payload", "message": "BAD REQUEST"})
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		log.Fatalf("Failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		log.Fatalf("Failed to parse end date: %v", err)
	}

	token, err := controller.HouseService.BuyHouseTransaction(ctx, userID, request.HouseID, startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	controller.EmailService.SendEmail(ctx, userEmail, "Konfirmasi pembelian sekarang", token.TransactionToken)

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

func (c *HouseController) GetHouseDetailWithTransactions(ctx echo.Context) error {
	houseID := ctx.Param("house_id")

	houseDetail, err := c.HouseService.GetHouseDetailWithTransactions(ctx.Request().Context(), houseID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch house details"})
	}

	return ctx.JSON(http.StatusOK, houseDetail)
}
