package mq

import (
	"context"

	"github.com/streadway/amqp"

	"grm-service/mq/message"
)

// 消息对象参数
type Options struct {
	Channel  *amqp.Channel
	Exchange []message.ExchangeOption
	Queue    []message.QueueOption
	Publish  []message.PublishOption
	Consume  []message.ConsumeOption

	// Other options
	Context context.Context
}
type Option func(*Options)

// Channel is the rabbitmq channel to user
func Channel(chn *amqp.Channel) Option {
	return func(o *Options) {
		o.Channel = chn
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func Publish(opt message.PublishOption) Option {
	return func(o *Options) {
		o.Publish = append(o.Publish, opt)
	}
}
