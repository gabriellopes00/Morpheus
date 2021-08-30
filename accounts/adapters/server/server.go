package server

import (
	"accounts/adapters/db"
	"accounts/adapters/queue"
	"accounts/adapters/server/handlers"
	"accounts/application/usecases"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

func SetupServer(router *echo.Echo, database *sql.DB, rabbitmq *amqp.Channel) {

	// init repository
	accountRepo := db.NewPgAccountRepository(database)

	// init queue
	rabbitMQ := queue.NewRabbitMQ(rabbitmq)

	// init usecases
	createAccount := usecases.NewAccountUsecase(accountRepo)

	// init handlers
	accountHandler := handlers.NewAccountHandler(createAccount, rabbitMQ)

	// setup middlewares
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())
	router.Use(middleware.Secure())
	router.Use(middleware.BodyLimit("2M"))
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// setup routes
	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "Ok"})
	})
	router.GET("/error", func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "Err"})
	})

	router.POST("/accounts", accountHandler.Create)
	router.POST("/signin", accountHandler.Auth)
	router.PUT("/accounts", accountHandler.Auth)
	router.GET("/accounts/:id", accountHandler.Create)
	router.DELETE("/accounts/:id", accountHandler.Create)
}
