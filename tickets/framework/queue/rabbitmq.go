package queue

import (
	"time"

	"github.com/streadway/amqp"
)

type MessageQueue interface {
	Consume(queue string, channel chan<- []byte)
	PublishMessage(exchange, key string, message []byte) error
}

type rabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitMQ(channel *amqp.Channel) *rabbitMQ {
	return &rabbitMQ{
		Channel: channel,
	}
}

func (rabbitmq *rabbitMQ) Consume(queue string, channel chan<- []byte) {
	msgs, err := rabbitmq.Channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for m := range msgs {
			channel <- m.Body
		}
		close(channel)
	}()
}

func (rabbitmq *rabbitMQ) PublishMessage(exchange, key string, message []byte) error {
	return rabbitmq.Channel.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			Timestamp:   time.Now().Local(),
		})
}
