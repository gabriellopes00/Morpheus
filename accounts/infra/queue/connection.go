package queue

import (
	"accounts/config/env"
	"accounts/interfaces"
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

	err = channel.ExchangeDeclare("accounts_ex", amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare("account_created", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare("account_deleted", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare("account_updated", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind("account_created", interfaces.QueueAccountCreated, "accounts_ex", false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind("account_deleted", interfaces.QueueAccountDeleted, "accounts_ex", false, nil)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind("account_updated", interfaces.QueueAccountUpdated, "accounts_ex", false, nil)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
