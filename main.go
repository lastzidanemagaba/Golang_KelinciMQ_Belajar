package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// help me connect rabbitmq
	conn, err := amqp.Dial("amqp://zidanemagaba:zidanemagaba@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// create a queue
	q, err := ch.QueueDeclare(
		"Antrian Saya", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// publish a message
	body := "Kamu Lagi Sama Doi Tapi Kamu Gak Tau"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [x] Sent %s", body)

	// consume a message
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
