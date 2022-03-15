package main

import (
	"context"
	"errors"
	"events/config/env"
	"events/framework/api"
	"events/framework/db"
	"events/framework/queue"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo/v4"
)

func main() {

	database, err := db.NewPostgresDb()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(database)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}

	defer database.Close()

	amqpConn, err := queue.NewRabbitMQConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer amqpConn.Close()

	queue.SetUpQueueServer(amqpConn, database)

	e := echo.New()
	api.SetupServer(e, database, amqpConn)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", env.PORT)); err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
