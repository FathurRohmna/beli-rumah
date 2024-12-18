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

	transactionID := c.Param("transaction_id")

	if err := h.UserHouseTransactionService.CancelTransaction(ctx, "", transactionID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction cancelled"})
}

func (h *UserHouseTransactionController) ConfirmTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	transactionID := c.Param("transaction_id")

	if err := h.UserHouseTransactionService.ConfirmTransaction(ctx, "", transactionID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction confirmed as sold"})
}
