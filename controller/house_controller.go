package controller

import "github.com/labstack/echo/v4"

type IHouseController interface {
	BuyHouseTransaction(c echo.Context) error
	GetHouses(c echo.Context) error
}
