package main

import (
	"context"
	"errors"
	"events/config/env"
	"events/framework/api"
	"events/framework/db"
	"events/framework/logger"
	"events/framework/queue"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {

	// postgres connection
	database, err := db.NewPostgresDb()
	if err != nil {
		logger.Logger.Fatal("error initializing database", zap.String("error_message", err.Error()))
	}

	sqlDb, err := database.DB()
	if err != nil {
		logger.Logger.Fatal("error getting sql db", zap.String("error_message", err.Error()))
	}

	defer sqlDb.Close()

	// migrations run
	err = db.AutoMigrate(sqlDb)
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Logger.Info("no pending migrations")
		} else {
			logger.Logger.Fatal("error running migrations", zap.String("error_message", err.Error()))
		}
	}

	// connecting to rabbitmq
	amqpConn, err := queue.NewRabbitMQConnection()
	if err != nil {
		logger.Logger.Fatal("error connecting to rabbitmq", zap.String("error_message", err.Error()))
	}

	defer amqpConn.Close()

	queue.SetUpQueueServer(amqpConn, database)

	e := echo.New()
	api.SetupServer(e, database, amqpConn)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", env.PORT)); err != nil {
			logger.Logger.Fatal("error starting web server", zap.String("error_message", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("error shutting down web server", zap.String("error_message", err.Error()))
	}
}
