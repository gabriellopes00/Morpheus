package queue

import (
	"accounts/config/env"
	"fmt"

	"github.com/streadway/amqp"
)

func NewRabbitMQConnection() (*amqp.Channel, error) {
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
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(ExchangeAccounts, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
