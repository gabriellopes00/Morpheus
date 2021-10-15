package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tickets/application"
	"tickets/framework/db"
	"tickets/framework/db/repositories"
	"tickets/framework/queue"
	"tickets/framework/queue/handlers"
	"time"

	"github.com/golang-migrate/migrate"
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

	rabbitmq := queue.NewRabbitMQ(amqpConn)

	channel := make(chan []byte)
	rabbitmq.Consume("event_created_tickets", channel)

	repo := repositories.NewPgEventsRepository(database)
	create := application.NewCreateEvent(repo)

	go func() {
		createHandler := handlers.NewCreateEventHandler(channel, create)
		err = createHandler.Create()
		if err != nil {
			panic(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
