package controller

import (
	"net/http"
	"beli-tanah/service"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService service.IPaymentService
}

func NewPaymentController(activityLogService service.IPaymentService) IPaymentController {
	return &PaymentController{
		PaymentService: activityLogService,
	}
}

func (controller *PaymentController) TopUpUserWallet(c echo.Context) error {
	ctx := c.Request().Context()
	posts := controller.PaymentService.TopUpUserWalletGeneratePayment(ctx, "f9f58304-a625-4037-9acc-e1bff70db0a0", 12000)

	return c.JSON(http.StatusOK, posts)
}
