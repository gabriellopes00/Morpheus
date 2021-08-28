package http

import (
	"accounts/application"
	"accounts/infra/crypto"
	"accounts/infra/db"
	"accounts/infra/http/handlers"
	"accounts/infra/queue"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupGinRouter(router *gin.Engine, connection *sql.DB) {
	createAccountHandler := handlers.NewCreateAccountHandler(
		application.NewCreateAccount(
			crypto.NewUUIDGenerator(),
			crypto.NewBcryptHasher(),
			db.NewPgAccountRepository(connection),
		),
		queue.NewRabbitMQ(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	router.POST("/accounts", createAccountHandler.Create)
}
