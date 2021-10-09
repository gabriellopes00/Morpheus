package main

import (
	"editions/config/env"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%d%s",
		env.RABBITMQ_USER,
		env.RABBITMQ_PASS,
		env.RABBITMQ_HOST,
		env.RABBITMQ_PORT,
		env.RABBITMQ_VHOST,
	)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare("event_created_edition", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind("event_created_edition", "event_created", "events_ex", false, nil)
	if err != nil {
		panic(err)
	}

	chann := make(chan []byte)
	msgs, err := channel.Consume("event_created_edition", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	go func() {
		for m := range msgs {
			chann <- m.Body
		}
		close(chann)
	}()

	for message := range chann {
		var data struct {
			Id string `json:"id,omitempty"`
		}
		err := json.Unmarshal(message, &data)
		if err != nil {
			panic(err)
		}

		fmt.Println(data)

	}
}
