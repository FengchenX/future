package main

import (
	"encoding/json"
	"sub_account_service/number_server/models"

	"flag"
	"sub_account_service/number_server/config"

	"github.com/nsqio/go-nsq"
	. "sub_account_service/number_server/pkg/nsq"
)

type MessageHandler struct {
	data chan *nsq.Message
}

func (m *MessageHandler) HandleMessage(message *nsq.Message) error {
	m.data <- message
	return nil
}

func (m *MessageHandler) Process() {
	for {
		message := <-m.data
		order := models.Orders{}
		json.Unmarshal([]byte(message.Body), &order)
		models.AddOrder(&order)
	}
}

func main() {

	flag.Parse()

	models.Setup()

	consumer := NewConsumer("order", "finance_order", []string{config.Opts().Nsqd_Consumer_Tcp}, []string{})

	handler := new(MessageHandler)

	handler.data = make(chan *nsq.Message, 1000)

	err := consumer.Start(nsq.HandlerFunc(handler.HandleMessage))

	if err != nil {
		panic(err)
	}

	//写入到数据库
	handler.Process()

}
