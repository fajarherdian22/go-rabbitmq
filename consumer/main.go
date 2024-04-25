package main

import (
	"log"

	"github.com/streadway/amqp"
)

func isError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	conRabbit, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	isError(err)
	defer conRabbit.Close()

	chanRabbit, err := conRabbit.Channel()

	isError(err)
	defer chanRabbit.Close()

	msg1, err := chanRabbit.Consume(
		"QueueService1",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	isError(err)

	msg2, err := chanRabbit.Consume(
		"QueueService2",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	isError(err)
	log.Println("Waiting messages")

	forever := make(chan bool)

	go func() {
		for message := range msg1 {
			log.Printf("Received messages from %s : %s \n", "QueueService1", message.Body)
		}
	}()

	go func() {
		for message := range msg2 {
			log.Printf("Received messages from %s : %s \n", "QueueService2", message.Body)
		}
	}()

	<-forever
}
