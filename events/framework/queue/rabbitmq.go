package queue

import (
	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitMQ(connection *amqp.Channel) *rabbitMQ {
	return &rabbitMQ{
		Channel: connection,
	}
}

func Consume(queue string, ch *amqp.Channel, in chan<- []byte) {

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)

	if err != nil {
		panic(err.Error())
	}

	msgs, err := ch.Consume(q.Name, "events", true, false, false, false, nil)

	if err != nil {
		panic(err.Error())
	}

	go func() {
		for m := range msgs {
			in <- []byte(m.Body)
		}
		close(in)
	}()
}
