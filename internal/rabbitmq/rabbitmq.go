package rabbitmq

import "github.com/streadway/amqp"

// Service -
type Service interface {
	Connect() error
	Publish(message string) error
}

// RabbitMQ - struct
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}
