package message

import (
	"log"
	"zidane/connection"

	"github.com/pkg/errors"
)

func Receive() {
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

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to consume queue"))
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To Exit press CTRL+C\n")
	<-forever
}
