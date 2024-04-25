package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func isError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Tutorial Rabbit MQ")
	conRabbit, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	isError(err)
	defer conRabbit.Close()

	chanRabbit, err := conRabbit.Channel()
	isError(err)
	defer chanRabbit.Close()

	q, err := chanRabbit.QueueDeclare(
		"QueueService1",
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Printf("this is Queue %v", q)

	isError(err)

	q2, err := chanRabbit.QueueDeclare(
		"QueueService2",
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Printf("this is Queue %v", q2)

	isError(err)

	app := gin.New()

	app.Use(gin.Logger())

	app.GET("/send", func(c *gin.Context) {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}
		if err := chanRabbit.Publish(
			"",
			q.Name,
			false,
			false,
			message,
		); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Waiting message for %v", q.Name)
	})
	app.GET("/spam2", func(c *gin.Context) {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}
		if err := chanRabbit.Publish(
			"",
			q2.Name,
			false,
			false,
			message,
		); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Waiting message for %v", q2.Name)
	})
	log.Fatal(app.Run(":3750"))
}
