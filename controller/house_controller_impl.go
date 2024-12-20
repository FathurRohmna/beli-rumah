package controller

import (
	"beli-tanah/exception"
	"beli-tanah/helper"
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
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error":   "unauthorized",
			"message": "User ID is invalid or missing. Please log in to continue.",
		})
	}

	userEmail, ok := c.Get("user_email").(string)
	if !ok || userEmail == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error":   "unauthorized",
			"message": "User email is invalid or missing. Please log in to continue.",
		})
	}

	ctx := c.Request().Context()
	var request web.BuyHouseTransactionRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "bad_request",
			"message": "Failed to parse request payload. Ensure the input format is correct.",
		})
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "bad_request",
			"message": "Start date format is invalid. Use 'YYYY-MM-DD'.",
		})
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		panic(exception.NewDataNotFoundError(fmt.Sprintf("Transaction error: %v", err)))
	}

	token, err := controller.HouseService.BuyHouseTransaction(ctx, userID, request.HouseID, startDate, endDate)
	if err != nil {
		panic(exception.NewDataNotFoundError(fmt.Sprintf("Transaction error: %v", err)))
	}

	data := map[string]string{
		"ExpiredDate": token.ExpiredAt.Local().Format("2 Jan 2006 15:04"),
		"MidtransUrl": token.TransactionToken,
		"Token":       token.TransactionToken,
	}

	emailBody, err := helper.RenderTemplate(data, "template/house_payment_confirmation.html")
	if err != nil {
		log.Fatalf("Error rendering template: %v", err)
	}

	controller.EmailService.SendEmail(ctx, userEmail, "Konfirmasi pembelian sekarang", emailBody)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Transaction successful. Confirmation email sent.",
	})
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
		houseCategory = ""
	}

	houses, totalCount, err := controller.HouseService.GetHouses(c.Request().Context(), houseCategory, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"houses":     houses,
		"totalCount": totalCount,
	})
}

func (c *HouseController) GetHouseDetailWithTransactions(ctx echo.Context) error {
	houseID := ctx.Param("houseId")

	houseDetail, err := c.HouseService.GetHouseDetailWithTransactions(ctx.Request().Context(), houseID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch house details"})
	}

	return ctx.JSON(http.StatusOK, houseDetail)
}
