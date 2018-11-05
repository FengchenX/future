package message

import (
	"context"
	"time"
)

type Table map[string]interface{}

// exchange 参数
type ExchangeOptions struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       Table

	// Other options
	Context context.Context
}
type ExchangeOption func(*ExchangeOptions)

func ExName(name string) ExchangeOption {
	return func(o *ExchangeOptions) {
		o.Name = name
	}
}

func ExKind(kind string) ExchangeOption {
	return func(o *ExchangeOptions) {
		o.Kind = kind
	}
}

func ExDurable(durable bool) ExchangeOption {
	return func(o *ExchangeOptions) {
		o.Durable = durable
	}
}

func ExAutoDelete(delete bool) ExchangeOption {
	return func(o *ExchangeOptions) {
		o.AutoDelete = delete
	}
}

func ExInternal(internal bool) ExchangeOption {
	return func(o *ExchangeOptions) {
		o.Internal = internal
	}
}

func ExNoWait(noWait bool) ExchangeOption {
	return func(o *ExchangeOptions) {
		o.NoWait = noWait
	}
}

// queue 参数
type QueueOptions struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       Table

	// Other options
	Context context.Context
}
type QueueOption func(*QueueOptions)

func QueueName(name string) QueueOption {
	return func(o *QueueOptions) {
		o.Name = name
	}
}

func QueueDurable(durable bool) QueueOption {
	return func(o *QueueOptions) {
		o.Durable = durable
	}
}

func QueueAutoDelete(delete bool) QueueOption {
	return func(o *QueueOptions) {
		o.AutoDelete = delete
	}
}

func QueueExclusive(exclusive bool) QueueOption {
	return func(o *QueueOptions) {
		o.Exclusive = exclusive
	}
}

func QueueNoWait(noWait bool) QueueOption {
	return func(o *QueueOptions) {
		o.NoWait = noWait
	}
}

// publish 参数
type PublishOptions struct {
	Exchange  string
	Queue     string
	Key       string
	Mandatory bool
	Immediate bool

	// MSG
	Headers         Table
	ContentType     string    // MIME content type
	ContentEncoding string    // MIME content encoding
	DeliveryMode    uint8     // Transient (0 or 1) or Persistent (2)
	Priority        uint8     // 0 to 9
	CorrelationId   string    // correlation identifier
	ReplyTo         string    // address to to reply to (ex: RPC)
	Expiration      string    // message expiration spec
	MessageId       string    // message identifier
	Timestamp       time.Time // message timestamp
	Type            string    // message type name
	UserId          string    // creating user id - ex: "guest"
	AppId           string    // creating application id
	// The application specific payload of the message
	Body []byte

	// Other options
	Context context.Context
}
type PublishOption func(*PublishOptions)

func PubExchange(name string) PublishOption {
	return func(o *PublishOptions) {
		o.Exchange = name
	}
}

func PubQueue(queue string) PublishOption {
	return func(o *PublishOptions) {
		o.Queue = queue
	}
}

func RoutKey(key string) PublishOption {
	return func(o *PublishOptions) {
		o.Key = key
	}
}

func PubBody(body []byte) PublishOption {
	return func(o *PublishOptions) {
		o.Body = body
	}
}

// consume 参数
type ConsumeOptions struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool

	// Other options
	Context context.Context
}
type ConsumeOption func(*ConsumeOptions)

func ConsumeQueue(queue string) ConsumeOption {
	return func(o *ConsumeOptions) {
		o.Queue = queue
	}
}
