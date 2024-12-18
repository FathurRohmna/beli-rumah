package controller

import "github.com/labstack/echo/v4"

type IUserHouseTransactionController interface {
	CancelTransactionHandler(c echo.Context) error
	ConfirmTransactionHandler(c echo.Context) error
}
