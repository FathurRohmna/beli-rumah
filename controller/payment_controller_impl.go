package controller

import (
	"beli-tanah/helper"
	"beli-tanah/model/domain"
	"beli-tanah/model/web"
	"beli-tanah/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService service.IPaymentService
	EmailService   service.IEmailService
	UserService    service.IUserService
}

func NewPaymentController(activityLogService service.IPaymentService, emailService service.IEmailService, userService service.IUserService) IPaymentController {
	return &PaymentController{
		PaymentService: activityLogService,
		EmailService:   emailService,
		UserService:    userService,
	}
}

func (controller *PaymentController) TopUpUserWallet(c echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing user ID", "message": "UNAUTHORIZED"})
	}

	var request web.TopUpUserWalletGeneratePaymentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"info": "Invalid request payload", "message": "BAD REQUEST"})
	}

	ctx := c.Request().Context()
	posts := controller.PaymentService.TopUpUserWalletGeneratePayment(ctx, userID, request.Amount)

	expirationTime := time.Now().Add(24 * time.Hour)
	data := map[string]string{
		"ExpiredDate": expirationTime.Format("2 Jan 2006 15:04"),
		"MidtransUrl": posts.PaymentUrl,
	}

	emailBody, err := helper.RenderTemplate(data, "template/request_top_up_success.html")
	if err != nil {
		log.Fatalf("Error rendering template: %v", err)
	}

	user := controller.UserService.GetUserById(ctx, userID)
	fmt.Print(user)
	err = controller.EmailService.SendEmail(ctx, user.Email, "Top Up success", emailBody)
	if err != nil {
		log.Fatalf("Error send email: %v", err)
	}

	return c.JSON(http.StatusOK, posts)
}

func (controller *PaymentController) MidtransCallback(c echo.Context) error {
	ctx := c.Request().Context()

	var callbackPayload domain.MidtransCallback
	if err := c.Bind(&callbackPayload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid payload",
		})
	}

	isValid := controller.PaymentService.VerifyMidtransSignature(callbackPayload)
	if !isValid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid signature",
		})
	}

	if callbackPayload.TransactionStatus == "settlement" || callbackPayload.TransactionStatus == "capture" {
		err := controller.PaymentService.UpdateWalletAndTransaction(ctx, callbackPayload.OrderID, callbackPayload.GrossAmount)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
	} else if callbackPayload.TransactionStatus == "cancel" || callbackPayload.TransactionStatus == "expire" || callbackPayload.TransactionStatus == "deny" {
		log.Printf("Transaction %s was not successful. Status: %s\n", callbackPayload.OrderID, callbackPayload.TransactionStatus)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Callback processed successfully",
	})
}
