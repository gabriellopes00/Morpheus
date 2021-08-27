package http

import (
	"accounts/application"
	"accounts/infra/crypto"
	"accounts/infra/db"
	"accounts/infra/http/handlers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupGinRouter(router *gin.Engine, connection *sql.DB) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	createAccountHandler := handlers.NewCreateAccountHandler(
		application.NewCreateAccount(
			crypto.NewUUIDGenerator(),
			crypto.NewBcryptHasher(),
			db.NewPgAccountRepository(connection),
		),
	)

	router.POST("/accounts", createAccountHandler.Create)
}
