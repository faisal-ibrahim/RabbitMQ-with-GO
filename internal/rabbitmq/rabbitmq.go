package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Service -struct
type Service interface {
	Connect() error
	Publish(message string) error
}

// RabbitMQ - struct
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Publish - publishes a message to the queue
func (r *RabbitMQ) Publish(message string) error {
	// attempt to publish a message to the queue!
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Successfully Published Message to Queue.")
	return nil
}

// NewRabbitMQService - returns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
