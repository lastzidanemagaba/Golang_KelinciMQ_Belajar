package message

import (
	"bufio"
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

	fmt.Println("Enter Your Message Below: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(scanner.Text()),
		})
		if err != nil {
			panic(errors.Wrap(err, "failed to publish message"))
		}
		fmt.Println("Message", scanner.Text()+" Sent Successfully")
		fmt.Println("Enter Your Message Again Below: ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
