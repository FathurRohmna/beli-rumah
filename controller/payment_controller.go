package controller

import "github.com/labstack/echo/v4"

type IPaymentController interface {
	TopUpUserWallet(c echo.Context) error
}
