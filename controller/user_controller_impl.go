package controller

import (
	"beli-tanah/model/web"
	"beli-tanah/service"
	"beli-tanah/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService   service.IUserService
	UserValidator validator.IUserValidator
}

func NewUserController(userService service.IUserService, userValidator validator.IUserValidator) *UserController {
	return &UserController{
		UserService:   userService,
		UserValidator: userValidator,
	}
}

func (controller *UserController) Register(c echo.Context) error {
	var userRequest web.RegisterUserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"info": "Invalid request payload", "message": "BAD REQUEST"})
	}

	if err := controller.UserValidator.ValidateRegisterUser(userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"info": err.Error(), "message": "BAD REQUEST"})
	}

	ctx := c.Request().Context()
	userResponse := controller.UserService.Register(ctx, userRequest)

	return c.JSON(http.StatusCreated, userResponse)
}

func (controller *UserController) Login(c echo.Context) error {
	var loginRequest web.LoginUserRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"info": "Invalid request payload", "message": "BAD REQUEST"})
	}

	if err := controller.UserValidator.ValidateLoginUser(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"info": err.Error(), "message": "BAD REQUEST"})
	}

	ctx := c.Request().Context()
	token := controller.UserService.Login(ctx, loginRequest)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (controller *UserController) GetMyDetail(c echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing user ID", "message": "UNAUTHORIZED"})
	}

	myDetail := controller.UserService.GetMyDetail(c.Request().Context(), userID)

	return c.JSON(http.StatusOK, myDetail)
}
