package rabbitmq

import "github.com/streadway/amqp"

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

// NewRabbitMQService - returns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
