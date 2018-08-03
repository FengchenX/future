package nsq

import (
	"errors"

	"github.com/nsqio/go-nsq"
)

//消息生产者
type Producer struct {
	producer *nsq.Producer
}

func NewProducer(addr string) (*Producer, error) {
	var p Producer
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		panic(err)
		return &p, err
	}
	p.producer = producer
	return &p, err
}

//发布一条message
func (p *Producer) Publish(topic string, body []byte) error {
	return p.producer.Publish(topic, body)
}

func (p *Producer) Stop() {
	p.producer.Stop()
}

//消费者
type Consumer struct {
	consumer    *nsq.Consumer
	config      *nsq.Config
	nsqds       []string
	nsqlookupds []string
	topic       string
	channel     string
	err         error
}

// NewConsumer returns a new consumer_server of `topic` and `channel`.
func NewConsumer(topic, channel string, nsqds []string, nsqlookupds []string) *Consumer {
	return &Consumer{
		config:      nsq.NewConfig(),
		topic:       topic,
		channel:     channel,
		nsqds:       nsqds,
		nsqlookupds: nsqlookupds,
	}
}

func (c *Consumer) Start(handle nsq.Handler) error {

	if c.err != nil {
		return c.err
	}

	consumer, err := nsq.NewConsumer(c.topic, c.channel, c.config)
	if err != nil {
		c.err = err
		return c.err
	}

	c.consumer = consumer

	c.consumer.AddHandler(handle)

	return c.connect()

}

func (c *Consumer) connect() error {

	if len(c.nsqds) == 0 && len(c.nsqlookupds) == 0 {
		c.err = errors.New(`at least one "nsqd" or "nsqlookupd" address must be configured`)
		return c.err
	}

	if len(c.nsqds) > 0 {
		err := c.consumer.ConnectToNSQDs(c.nsqds)
		if err != nil {
			c.err = err
			return c.err
		}
	}

	if len(c.nsqlookupds) > 0 {
		err := c.consumer.ConnectToNSQLookupds(c.nsqlookupds)
		if err != nil {
			c.err = err
			return c.err
		}
	}
	return nil
}
