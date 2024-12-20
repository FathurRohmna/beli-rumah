package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetMyDetail(c echo.Context) error
}
