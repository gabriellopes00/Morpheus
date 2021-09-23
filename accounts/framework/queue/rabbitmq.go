package queue

import (
	"time"

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

func (r *rabbitMQ) SendMessage(queue string, payload []byte) error {
	err := r.Channel.Publish(
		"accounts_ex",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Timestamp:   time.Now().Local(),
		})

	return err
}
