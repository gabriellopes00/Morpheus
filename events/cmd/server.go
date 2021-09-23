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

	in_create := make(chan []byte)
	in_delete := make(chan []byte)

	queue.Consume("account_created_events", rabbitmq, in_create)
	queue.Consume("account_deleted_events", rabbitmq, in_delete)

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
}
