package server

import (
	"accounts/adapters/db"
	"accounts/adapters/queue"
	"accounts/adapters/server/handlers"
	"accounts/application/usecases"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
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

	// setup routes
	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "Ok"})
	})

	router.POST("/accounts", accountHandler.Create)
	router.POST("/signin", accountHandler.Auth)
}
