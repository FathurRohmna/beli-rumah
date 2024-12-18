package controller

import (
	"beli-tanah/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HouseController struct {
	HouseService service.IHouseService
}

func NewHouseController(houseService service.IHouseService) IHouseController {
	return &HouseController{HouseService: houseService}
}

func (h *HouseController) BuyHouseTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.HouseService.BuyHouseTransaction(ctx, "", ""); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction cancelled"})
}
