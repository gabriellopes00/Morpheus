package main

import (
	"accounts/config/env"
	"accounts/framework/api"
	"accounts/framework/db"
	"accounts/framework/queue"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}

	defer logger.Sync()

	database, err := db.NewPostgresDb()
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = db.AutoMigrate(database)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := database.Close(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	rabbitmq, err := queue.NewRabbitMQConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := rabbitmq.Close(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	e := echo.New()
	api.SetupServer(e, database, rabbitmq)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", env.PORT)); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
	}
}
