package pkgmiddleware

import (
	"beli-tanah/service"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func TransactionTokenMiddleware(userService service.IUserService, transactionService service.IUserHouseTransactionService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.QueryParam("tokenId")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Token is required", "message": "UNAUTHORIZED"})
			}

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_AUTH_EMAIL_URL")), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid token", "message": "UNAUTHORIZED"})
			}

			userID, ok := claims["user_id"].(string)
			if !ok || userID == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid token: user ID not found", "message": "UNAUTHORIZED"})
			}

			transactionID, ok := claims["transaction_id"].(string)
			if !ok || transactionID == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid token: transaction ID not found", "message": "UNAUTHORIZED"})
			}

			ctx := c.Request().Context()
			user := userService.GetUserById(ctx, userID)
			transaction, err := transactionService.FindTransactionById(ctx, transactionID)

			if err != nil || user.ID != transaction.UserID {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid authorization", "message": "UNAUTHORIZED"})
			}

			if transaction.Status != "pending" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Transaction is not pending", "message": "UNAUTHORIZED"})
			}

			now := time.Now()
			if transaction.ExpiredAt.Before(now) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Transaction expired", "message": "UNAUTHORIZED"})
			}

			c.Set("user_id", userID)
			c.Set("transaction_id", transactionID)
			c.Set("house_id", transaction.HouseID)
			return next(c)
		}
	}
}

func HouseAvailabilityMiddleware(houseService service.IHouseService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			houseID, ok := c.Get("house_id").(string)
			if !ok || houseID == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "Invalid or missing house ID", "message": "UNAUTHORIZED"})
			}

			ctx := c.Request().Context()

			if err := houseService.CheckHouseAvailability(ctx, houseID); err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"info": "House is not available", "message": "UNAUTHORIZED"})
			}

			return next(c)
		}
	}
}
