package main

import (
	"events/application"
	"events/framework/db"
	"events/framework/db/repositories"
	"events/framework/queue"
	"events/framework/queue/handlers"
	"fmt"
	"log"
	"os"
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

	defer func() {
		if err := database.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	rabbitmq, err := queue.NewRabbitMQConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := rabbitmq.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	in := make(chan []byte)

	queue.Consume("account_created", rabbitmq, in)

	repo := repositories.NewPgAccountRepository(database)
	u := application.CreateAccountUsecase{Repository: repo}

	h := handlers.NewCreateAccountHandler(in, u)
	err = h.Create()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("app finished")
	os.Exit(0)
}
