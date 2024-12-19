package controller

import (
	"beli-tanah/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHouseTransactionController struct {
	UserHouseTransactionService service.IUserHouseTransactionService
}

func NewUserHouseTransactionController(userHouseTransactionService service.IUserHouseTransactionService) IUserHouseTransactionController {
	return &UserHouseTransactionController{UserHouseTransactionService: userHouseTransactionService}
}

func (h *UserHouseTransactionController) CancelTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()
	transactionID, ok := c.Get("transaction_id").(string)
	if !ok || transactionID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing house ID", "message": "UNAUTHORIZED"})
	}

	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing user ID", "message": "UNAUTHORIZED"})
	}

	if err := h.UserHouseTransactionService.CancelTransaction(ctx, userID, transactionID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction cancelled"})
}

func (h *UserHouseTransactionController) ConfirmTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()
	transactionID, ok := c.Get("transaction_id").(string)
	if !ok || transactionID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing house ID", "message": "UNAUTHORIZED"})
	}

	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing user ID", "message": "UNAUTHORIZED"})
	}

	if err := h.UserHouseTransactionService.ConfirmTransaction(ctx, userID, transactionID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction confirmed as sold"})
}
