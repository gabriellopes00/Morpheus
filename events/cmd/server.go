package main

import (
	"events/application"
	"events/framework/db"
	"events/framework/db/repositories"
	"events/framework/queue"
	"events/framework/queue/handlers"
	"fmt"
	"log"
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

	rabbitmq, err := queue.NewRabbitMQConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer rabbitmq.Close()

	in := make(chan []byte)

	queue.Consume("account_created", rabbitmq, in)

	repo := repositories.NewPgAccountRepository(database)
	u := application.CreateAccountUsecase{Repository: repo}

	h := handlers.NewCreateAccountHandler(in, u)
	err = h.Create()
	if err != nil {
		panic(err.Error())
	}
}
