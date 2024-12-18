package pkgmiddleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Token is required", "message": "UNAUTHORIZED"})
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid token", "message": "UNAUTHORIZED"})
		}

		studentID, ok := claims["user_id"].(string)
		if !ok || studentID == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid token: user ID not found", "message": "UNAUTHORIZED"})
		}

		c.Set("user_id", studentID)
		return next(c)
	}
}
