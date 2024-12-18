package controller

import (
	"beli-tanah/service"
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService service.IPaymentService
	EmailService   service.IEmailService
}

func RenderTemplate(data map[string]string) (string, error) {
	tmpl, err := template.ParseFiles("template/email_template.html")
	if err != nil {
		return "", err
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, data); err != nil {
		return "", err
	}

	return rendered.String(), nil
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

	data := map[string]string{
		"Name":  "Fathur",
		"Event": "meeting",
		"Date":  "Monday, Dec 25th",
	}

	emailBody, err := RenderTemplate(data)
	if err != nil {
		log.Fatalf("Error rendering template: %v", err)
	}

	controller.EmailService.SendEmail(ctx, "fatur23460@gmail.com", "Testing email here", emailBody)

	return c.JSON(http.StatusOK, posts)
}
