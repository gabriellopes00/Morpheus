package main

import (
	"accounts/infra/db"
	"accounts/infra/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	postgres := db.NewPostgresDb()

	connection, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer connection.Close()

	router := gin.Default()
	http.SetupGinRouter(router, connection)
	router.Run(":7765")
}
