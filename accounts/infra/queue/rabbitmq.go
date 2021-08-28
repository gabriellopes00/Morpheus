package queue

import (
	env "accounts/config"
	"fmt"

	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitMQ() *rabbitMQ {
	return &rabbitMQ{}
}

func (r *rabbitMQ) Connect() error {
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
		return err
	}

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	r.Channel = channel

	return nil
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
		})

	return err
}
