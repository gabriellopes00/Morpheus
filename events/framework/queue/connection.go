package queue

import (
	"events/config/env"
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

	_, err = channel.QueueDeclare("account_created_events", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare("account_deleted_events", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind("account_created_events", "account_created", "accounts_ex", false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind("account_deleted_events", "account_deleted", "accounts_ex", false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare("events_ex", amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare("event_created", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind("event_created", "event_created", "events_ex", false, nil)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
