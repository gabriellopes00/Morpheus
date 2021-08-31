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

func (r *rabbitMQ) SendMessage(payload []byte) error {
	err := r.Channel.Publish(
		"accounts_ex",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
			Timestamp:   time.Now().Local(),
		})

	return err
}
