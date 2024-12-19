package controller

import "github.com/labstack/echo/v4"

type IPaymentController interface {
	TopUpUserWallet(c echo.Context) error
	MidtransCallback(c echo.Context) error 
}
