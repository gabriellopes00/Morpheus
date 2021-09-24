package main

import (
	"context"
	"events/application"
	"events/config/env"
	"events/framework/api"
	"events/framework/db"
	"events/framework/db/repositories"
	"events/framework/queue"
	"events/framework/queue/handlers"
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
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(database)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer database.Close()

	amqpConn, err := queue.NewRabbitMQConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer amqpConn.Close()

	rabbitmq := queue.NewRabbitMQ(amqpConn)

	in_create := make(chan []byte)
	in_delete := make(chan []byte)

	rabbitmq.Consume("account_created_events", in_create)
	rabbitmq.Consume("account_deleted_events", in_delete)

	repo := repositories.NewPgAccountRepository(database)
	create := application.CreateAccountUsecase{Repository: repo}
	delete := application.DeleteAccountUsecase{Repository: repo}

	go func() {
		createHandler := handlers.NewCreateAccountHandler(in_create, create)
		err = createHandler.Create()
		if err != nil {
			panic(err.Error())
		}
	}()

	deleteHandler := handlers.NewDeleteAccountHandler(in_delete, delete)
	err = deleteHandler.Delete()
	if err != nil {
		panic(err.Error())
	}

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
