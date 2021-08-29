package http

import (
	"accounts/adapters/crypto"
	"accounts/adapters/db"
	"accounts/adapters/http/handlers"
	"accounts/adapters/queue"
	"accounts/application/usecases"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func SetupRouter(router *gin.Engine, dbConn *sql.DB, queueConn *amqp.Channel) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	createAccount := usecases.NewCreateAccount(
		crypto.NewUUIDGenerator(),
		crypto.NewBcryptHasher(),
		db.NewPgAccountRepository(dbConn),
	)

	queue := queue.NewRabbitMQ(queueConn)

	createAccountHandler := handlers.NewCreateAccountHandler(createAccount, queue)
	router.POST("/accounts", createAccountHandler.Create)
}
