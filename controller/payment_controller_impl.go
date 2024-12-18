package controller

import (
	"beli-tanah/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService service.IPaymentService
	EmailService   service.IEmailService
}

func NewPaymentController(activityLogService service.IPaymentService, emailService service.IEmailService) IPaymentController {
	return &PaymentController{
		PaymentService: activityLogService,
		EmailService:   emailService,
	}
}

func (controller *PaymentController) TopUpUserWallet(c echo.Context) error {
	ctx := c.Request().Context()
	posts := controller.PaymentService.TopUpUserWalletGeneratePayment(ctx, "f9f58304-a625-4037-9acc-e1bff70db0a0", 12000)
	controller.EmailService.SendEmail(ctx, "fatur23460@gmail.com", "Testing email here", posts.PaymentUrl)

	return c.JSON(http.StatusOK, posts)
}
