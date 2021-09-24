package queue

import (
	"time"

	"github.com/streadway/amqp"
)

const (
	QueueEventCreated = "event_created"
)

type MessageQueue interface {
	Consume(queue string, channel chan<- []byte)
	PublishMessage(exchange, queue string, message []byte) error
}

type rabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitMQ(connection *amqp.Channel) *rabbitMQ {
	return &rabbitMQ{
		Channel: connection,
	}
}

func (rabbitmq *rabbitMQ) Consume(queue string, channel chan<- []byte) {

	q, err := rabbitmq.Channel.QueueDeclare(queue, true, false, false, false, nil)

	if err != nil {
		panic(err.Error())
	}

	msgs, err := rabbitmq.Channel.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		panic(err.Error())
	}

	go func() {
		for m := range msgs {
			channel <- []byte(m.Body)
		}
		close(channel)
	}()
}

func (rabbitmq *rabbitMQ) PublishMessage(queue, exchange string, payload []byte) error {
	return rabbitmq.Channel.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Timestamp:   time.Now().Local(),
		})
}
