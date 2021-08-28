package queue

import (
	env "accounts/config"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitMQ() *rabbitMQ {
	channel, err := connect()
	if err != nil {
		panic(err.Error())
	}

	return &rabbitMQ{
		Channel: channel,
	}
}

func connect() (*amqp.Channel, error) {
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

	err = channel.QueueBind("account_created", "", "accounts_ex", false, nil)
	if err != nil {
		return nil, err
	}

	return channel, nil
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
