package main

import (
	"accounts/config/env"
	"accounts/pkg/api"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"accounts/pkg/logger"
	"accounts/pkg/queue"
	"accounts/pkg/storage"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	database, err := db.NewPostgresDb()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	err = db.AutoMigrate(database)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := database.Close(); err != nil {
			logger.Logger.Fatal(err.Error())
		}
	}()

	rabbitmq, err := queue.NewRabbitMQConnection()
	if err != nil {
		fmt.Println("rabbitmq")
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := rabbitmq.Close(); err != nil {
			logger.Logger.Fatal(err.Error())
		}
	}()

	redis, err := cache.NewRedisClient()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	defer redis.Close()

	awsSession, err := storage.CreateAWSSession()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	e := echo.New()
	api.SetupServer(e, database, rabbitmq, redis, awsSession)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", env.PORT)); err != nil {
			logger.Logger.Fatal(err.Error())
		}
	}()

	// gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Logger.Fatal(err.Error())
	}
}
