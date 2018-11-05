package message

import (
	"github.com/streadway/amqp"
)

type Message interface {
	ExchangeDeclare(options ...ExchangeOption) error
	QueueDeclare(options ...QueueOption) (string, error)
	Publish(options ...PublishOption) error
	Consume(options ...ConsumeOption) (<-chan amqp.Delivery, error)
	CloseChannel() error
}
