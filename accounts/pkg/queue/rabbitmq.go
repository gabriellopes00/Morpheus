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

func (r *rabbitMQ) PublishMessage(exchange, key string, payload []byte) error {
	return r.Channel.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Timestamp:   time.Now().Local(),
		})
}
