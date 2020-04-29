package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)


func Rec(qName string) []byte {
	// подключение и объявление очереди вынести в отдельную функцию и не вызывать каждый раз
	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	failOnError(err, "Failed to connect to RabbitMQ")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		true,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var answ []byte
	fmt.Println("Wait for msg from", qName)

	// написать что-то, чтобы функция возвращала результат, даже если очередь пустая.
	// noWait стоит true, но как сделать, чтобы читалось сразу - не понимаю

	for d := range msgs {
		answ = d.Body
		break
	}
	fmt.Println("REC answer: ", string(answ))

	return answ
}
