package queue

import (
	"database/sql"
	"events/application"
	"events/framework/db/repositories"
	"events/framework/queue/handlers"

	"github.com/streadway/amqp"
)

func SetUpQueueServer(amqpConn *amqp.Channel, database *sql.DB) {
	rabbitmq := NewRabbitMQ(amqpConn)

	accountCreatedChan := make(chan []byte)
	accountDeletedChan := make(chan []byte)
	soldOutChan := make(chan []byte)

	rabbitmq.Consume("account_created_events", accountCreatedChan)
	rabbitmq.Consume("account_deleted_events", accountDeletedChan)
	rabbitmq.Consume("event_sold_out", soldOutChan)

	accountsRepository := repositories.NewPgAccountRepository(database)
	eventsRepository := repositories.NewPgEventsRepository(database)
	create := application.CreateAccountUsecase{Repository: accountsRepository}
	delete := application.DeleteAccountUsecase{Repository: accountsRepository}
	updateEvent := application.NewUpdateEventUsecase(eventsRepository)

	go func() {
		createHandler := handlers.NewCreateAccountHandler(accountCreatedChan, create)
		if err := createHandler.Create(); err != nil {
			panic(err.Error())
		}
	}()

	go func() {
		deleteHandler := handlers.NewDeleteAccountHandler(accountDeletedChan, delete)
		if err := deleteHandler.Delete(); err != nil {
			panic(err.Error())
		}
	}()

	go func() {
		soldOutEventHandler := handlers.NewsoldOutEventHandler(soldOutChan, updateEvent)
		if err := soldOutEventHandler.Handle(); err != nil {
			panic(err.Error())
		}
	}()
}
