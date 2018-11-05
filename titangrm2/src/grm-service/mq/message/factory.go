package message

import (
	"github.com/streadway/amqp"
)

type MsgFactory interface {
	NewMessage(*amqp.Channel) (Message, error)
}
