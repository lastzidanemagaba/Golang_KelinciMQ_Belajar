package connection

import (
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

const (
	// RabbitMQURL is the URL to connect to RabbitMQ
	RabbitMQURL = "amqp://zidanemagaba:zidanemagaba@localhost:5672/"
	Queuenya    = "AdaApaDenganCinta"
)

func RabbitMQ() (*amqp091.Connection, *amqp091.Channel, error) {
	conn, err := amqp091.Dial(RabbitMQURL)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get channel")
	}
	return conn, ch, nil
}
