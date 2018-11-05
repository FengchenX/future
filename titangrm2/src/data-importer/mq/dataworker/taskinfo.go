package dataworker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"

	"data-importer/dbcentral/etcd"
	"grm-service/mq"
	"grm-service/mq/message"
)

const (
	TASKMESSAGE = "TaskInfo"
	EXCHANGE    = "grm.data-importer.task-info"
	ROUTINGKEY  = ""
	EXTYPE      = "fanout"
)

type TaskInfoFactory struct {
}

// 注册消息工厂
func init() {
	mq.Register("TaskInfo", &TaskInfoFactory{})
}

func (f *TaskInfoFactory) NewMessage(chn *amqp.Channel) (message.Message, error) {
	msg := new(TaskInfoMsg)
	msg.channel = chn
	return msg, nil
}

type TaskInfoMsg struct {
	channel *amqp.Channel

	queue string
}

func PublishHandler(dynamic *etcd.DynamicDB, taskType, taskId string) message.PublishOption {
	return func(o *message.PublishOptions) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "dynamic_etcd", dynamic)
		ctx = context.WithValue(ctx, "task_type", taskType)
		ctx = context.WithValue(ctx, "task_id", taskId)
		o.Context = ctx
	}
}

func (m *TaskInfoMsg) ExchangeDeclare(options ...message.ExchangeOption) error {
	var option message.ExchangeOptions
	for _, o := range options {
		o(&option)
	}

	// exchange
	option.Name = EXCHANGE

	if len(option.Kind) == 0 {
		option.Kind = EXTYPE
	}

	return m.channel.ExchangeDeclare(
		option.Name, // name
		option.Kind, // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
}

func (m *TaskInfoMsg) QueueDeclare(options ...message.QueueOption) (string, error) {
	var option message.QueueOptions
	for _, o := range options {
		o(&option)
	}

	q, err := m.channel.QueueDeclare(
		option.Name, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return "", err
	}
	return q.Name, nil
}

func (m *TaskInfoMsg) Publish(options ...message.PublishOption) error {
	var option message.PublishOptions
	for _, o := range options {
		o(&option)
	}
	dynamic, ok := option.Context.Value("dynamic_etcd").(*etcd.DynamicDB)
	if !ok {
		return fmt.Errorf("Failed to get dynamic etcd connection")
	}
	taskType, ok := option.Context.Value("task_type").(string)
	if !ok {
		return fmt.Errorf("Failed to get task_type")
	}

	taskId, ok := option.Context.Value("task_id").(string)
	if !ok {
		return fmt.Errorf("Failed to get task_id")
	}
	info, err := dynamic.GetTaskInfo(taskType, taskId)
	if err != nil {
		return err
	}
	body, err := json.Marshal(info)
	if err != nil {
		return err
	}
	// exchange
	option.Exchange = EXCHANGE

	// routing key
	if len(option.Key) == 0 {
		option.Key = ROUTINGKEY
	}

	return m.channel.Publish(
		option.Exchange, option.Key, option.Mandatory, option.Immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Timestamp:   time.Now(),
			Body:        body,
		})
}

func (m *TaskInfoMsg) Consume(options ...message.ConsumeOption) (<-chan amqp.Delivery, error) {
	var option message.ConsumeOptions
	for _, o := range options {
		o(&option)
	}

	if err := m.channel.QueueBind(
		option.Queue, // queue name
		"",           // routing key
		EXCHANGE,     // exchange
		false,
		nil); err != nil {
		return nil, err
	}

	return m.channel.Consume(
		option.Queue, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
}

func (m *TaskInfoMsg) CloseChannel() error {
	return m.channel.Close()
}
