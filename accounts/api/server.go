package api

import (
	"accounts/api/handlers"
	"accounts/infra/db"
	"accounts/infra/encrypter"
	"accounts/infra/queue"
	"accounts/usecases"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

func SetupServer(router *echo.Echo, database *sql.DB, rabbitmq *amqp.Channel) {

	// init adapters
	accountRepo := db.NewPgAccountRepository(database)
	jwtEncrypter := encrypter.NewJwtEncrypter()
	rabbitMQ := queue.NewRabbitMQ(rabbitmq)

	// init usecases
	createAccount := usecases.NewCreateAccount(accountRepo)
	authAccount := usecases.NewAuthAccount(accountRepo, jwtEncrypter)

	// init handlers
	createAccountHandler := handlers.NewCreateAccountHandler(createAccount, rabbitMQ, jwtEncrypter)
	authHandler := handlers.NewAuthHandler(authAccount)

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

	router.POST("/accounts", createAccountHandler.Create)
	router.POST("/signin", authHandler.Auth)
	router.PUT("/accounts", createAccountHandler.Create)
	router.GET("/accounts/:id", createAccountHandler.Create)
	router.DELETE("/accounts/:id", createAccountHandler.Create)
}
