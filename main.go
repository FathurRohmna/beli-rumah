package main

import (
	"beli-tanah/app"
	"beli-tanah/controller"
	pkgmiddleware "beli-tanah/middleware"
	"beli-tanah/repository"
	"beli-tanah/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := app.NewDB()

	paymentRepository := repository.NewPaymentRepository()
	houseRepository := repository.NewHouseRepository()
	userHouseTransactionRepository := repository.NewUserHouseTransactionRepository()
	userRepository := repository.NewUserRepository()

	emailService := service.NewEmailService()
	paymentService := service.NewPaymentService(paymentRepository, db)
	houseService := service.NewHouseService(houseRepository, userHouseTransactionRepository, db)
	userHouseTransactionService := service.NewUserHouseTransactionService(userHouseTransactionRepository, db)
	userService := service.NewUserService(userRepository, db)

	paymentController := controller.NewPaymentController(paymentService, emailService)
	houseController := controller.NewHouseController(houseService, emailService)
	userHouseTransactionController := controller.NewUserHouseTransactionController(userHouseTransactionService)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Recover())

	e.POST("/topup", paymentController.TopUpUserWallet)

	e.POST("/buyhouse", houseController.BuyHouseTransaction)

	e.POST(
		"/transaction/cancel",
		userHouseTransactionController.CancelTransactionHandler,
		pkgmiddleware.TransactionTokenMiddleware(userService, userHouseTransactionService),
	)
	e.POST(
		"/transaction/confirm",
		userHouseTransactionController.ConfirmTransactionHandler,
		pkgmiddleware.TransactionTokenMiddleware(userService, userHouseTransactionService),
		pkgmiddleware.HouseAvailabilityMiddleware(houseService),
	)

	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
