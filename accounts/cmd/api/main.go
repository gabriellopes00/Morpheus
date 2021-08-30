package main

import (
	"accounts/adapters/db"
	"accounts/adapters/queue"
	"accounts/adapters/server"
	"accounts/config/env"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	database, err := db.NewPostgresDb()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	rabbitmq, err := queue.NewRabbitMQConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := rabbitmq.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	e := echo.New()
	server.SetupServer(e, database, rabbitmq)
	if err := e.Start(fmt.Sprintf(":%d", env.PORT)); err != nil {
		log.Fatalln(err)
	}
}
