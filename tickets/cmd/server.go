package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tickets/application"
	"tickets/framework/queue"
	"tickets/framework/queue/handlers"
	"time"
)

func main() {

	// database, err := db.NewPostgresDb()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// err = db.AutoMigrate(database)
	// if err != nil && !errors.Is(err, migrate.ErrNoChange) {
	// 	log.Fatalln(err)
	// }

	// defer database.Close()

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

	go func() {
		deleteHandler := handlers.NewDeleteAccountHandler(in_delete, delete)
		err = deleteHandler.Delete()
		if err != nil {
			panic(err.Error())
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
