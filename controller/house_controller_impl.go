package controller

import (
	"beli-tanah/service"
	"net/http"

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

	token, err := controller.HouseService.BuyHouseTransaction(ctx, "0e4f279d-03ea-46dc-a07a-92d057e1e470", "792d35f3-92cb-4c45-bad0-2042ab02c4aa")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	controller.EmailService.SendEmail(ctx, "fatur23460@gmail.com", "Testing email here", token.TransactionToken)

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction cancelled"})
}
