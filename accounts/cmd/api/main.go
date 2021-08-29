package main

import (
	"accounts/adapters/db"
	"accounts/adapters/http"
	"accounts/adapters/queue"
	"accounts/config/env"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	postgres := db.NewPostgresDb()
	dbConn, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer dbConn.Close()

	queueConn, err := queue.NewQueueConnection().Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.Default()
	http.SetupRouter(router, dbConn, queueConn)
	router.Run(fmt.Sprintf(":%d", env.PORT))
}
