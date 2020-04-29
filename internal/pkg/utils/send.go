package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send(dater []byte, qName string) {
	// подключение и объявление очереди вынести в отдельную функцию
	var conn *amqp.Connection
	for {
		conn2, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err == nil {
			conn = conn2
			fmt.Println("connected to RABBIT")
			break
		}
		fmt.Println("Not connected: ", err)
		time.Sleep(500 * time.Millisecond)
	}
	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		qName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dater,
		})
	log.Printf(" [x] Sent %s to %s", dater, qName)
	failOnError(err, "Failed to publish a message")
}
