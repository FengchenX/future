package mq

import (
	"fmt"
	"sort"
	"sync"

	"grm-service/log"
	"grm-service/mq/message"

	"github.com/streadway/amqp"
)

var (
	mutex     sync.RWMutex
	factories = make(map[string]message.MsgFactory)
)

func Register(msgType string, factory message.MsgFactory) {
	mutex.Lock()
	defer mutex.Unlock()
	if factory == nil {
		panic("mq: Message factory is nil")
	}
	if _, dup := factories[msgType]; dup {
		panic("mq: Register called twice for factory " + msgType)
	}
	factories[msgType] = factory
}

func MsgFactories() []string {
	mutex.RLock()
	defer mutex.RUnlock()
	var list []string
	for name := range factories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

// RabbitMQ对象
type RabbitMQ struct {
	URL         string
	VirtualHost string // TODO: 区分不同服务或者项目
	conn        *amqp.Connection
}

func (mq *RabbitMQ) Connect() error {
	conn, err := amqp.Dial(mq.URL)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ: ", mq.URL, err)
		return err
	}
	mq.conn = conn
	return nil
}

func (mq *RabbitMQ) Close() error {
	if mq != nil && mq.conn != nil {
		return mq.conn.Close()
	}
	return nil
}

func (mq *RabbitMQ) GetChannel() (*amqp.Channel, error) {
	return mq.conn.Channel()
}

func (mq *RabbitMQ) newMessage(msgType string, opts ...Option) (message.Message, string, error) {
	f, ok := factories[msgType]
	if !ok {
		return nil, "", fmt.Errorf("mq: no factory for message: %s", msgType)
	}

	// parse options
	var options Options
	for _, o := range opts {
		o(&options)
	}

	// channel
	var channel *amqp.Channel
	if options.Channel != nil {
		channel = options.Channel
	} else {
		chn, err := mq.GetChannel()
		if err != nil {
			return nil, "", err
		}
		channel = chn
	}

	msg, err := f.NewMessage(channel)
	if err != nil {
		return nil, "", err
	}

	// exchange
	if err := msg.ExchangeDeclare(options.Exchange...); err != nil {
		return nil, "", err
	}

	// queue
	name, err := msg.QueueDeclare(options.Queue...)
	if err != nil {
		return nil, "", err
	}
	return msg, name, nil
}

// 发送消息
func (mq *RabbitMQ) PublishMessage(msgType string, opts ...Option) error {
	msg, queue, err := mq.newMessage(msgType, opts...)
	if err != nil {
		return err
	}
	defer msg.CloseChannel()

	// parse options
	var options Options
	for _, o := range opts {
		o(&options)
	}
	options.Publish = append(options.Publish, message.PubQueue(queue))
	return msg.Publish(options.Publish...)
}

// 接收消息
func (mq *RabbitMQ) ConsumeMessage(msgType string, opts ...Option) (<-chan amqp.Delivery, error) {
	msg, queue, err := mq.newMessage(msgType, opts...)
	if err != nil {
		return nil, err
	}
	// parse options
	var options Options
	for _, o := range opts {
		o(&options)
	}
	options.Consume = append(options.Consume, message.ConsumeQueue(queue))
	return msg.Consume(options.Consume...)
}
