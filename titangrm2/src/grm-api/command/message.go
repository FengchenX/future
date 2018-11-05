package command

import (
	"fmt"

	"data-importer/mq/dataworker"
)

func (c *APICommand) SubScribeTaskInfo(hub *Hub) {
	go func() {
		msgs, err := c.MsgQueue.ConsumeMessage(dataworker.TASKMESSAGE)
		if err != nil {
			return
		}
		for d := range msgs {
			fmt.Println("Receive msg, time:", d.Timestamp, "body: ", string(d.Body))
			for k, v := range hub.clients {
				if v && k.clientType == dataworker.TASKMESSAGE {
					fmt.Println("Send websocket, time:", d.Timestamp, "body: ", string(d.Body))
					k.writeTaskInfo(d.Body)
				}
			}
		}
	}()
}
