package message

import (
	"fmt"
	"os"
	"zidane/connection"

	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

func Send() {
	conn, ch, err := connection.RabbitMQ()
	if err != nil {
		panic(err)
	}

	defer func() {
		ch.Close()
		conn.Close()
	}()

	q, err := ch.QueueDeclare(connection.Queuenya, false, false, false, false, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to declare queue"))
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(os.Args[2]),
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to publish message"))
	}

	fmt.Println("Send message:", os.Args[2])
}
