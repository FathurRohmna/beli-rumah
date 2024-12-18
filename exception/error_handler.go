package exception

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if authenticationCredentialError(err, c) {
		return
	}

	if dataNotFoundError(err, c) {
		return
	}

	internalServerError(err, c)
}

func authenticationCredentialError(err error, c echo.Context) bool {
	_, ok := err.(AuthenticationCredentialError)
	if ok {
		response := map[string]interface{}{
			"message": "INVALID CREDENTIAL",
			"info":    err.Error(),
		}

		c.JSON(http.StatusUnauthorized, response)
		return true
	} else {
		return false
	}
}

func dataNotFoundError(err error, c echo.Context) bool {
	_, ok := err.(DataNotFoundError)
	if ok {
		response := map[string]interface{}{
			"message": "DATA NOT FOUND",
			"info":    err.Error(),
		}

		c.JSON(http.StatusNotFound, response)
		return true
	} else {
		return false
	}
}

func internalServerError(err error, c echo.Context) {
	response := map[string]interface{}{
		"message": "INTERNAL SERVER ERROR",
		"info":    err.Error(),
	}

	c.JSON(http.StatusInternalServerError, response)
}
