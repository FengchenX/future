package command

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"

	. "grm-api/api"
	"grm-service/command"
	"grm-service/mq"
	"grm-service/service"
)

type APICommand struct {
	command.Meta

	MqUrl    string
	MsgQueue *mq.RabbitMQ
}

func (c *APICommand) Help() string {
	helpText := `
Usage: titan-grm grm-api [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm api -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *APICommand) Synopsis() string {
	return "GRM API Gateway"
}

func (c *APICommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.GRMAPIService, command.FlagSetDefault)

	flags.StringVar(&c.MqUrl, "mq", "amqp://admin:otitan123@192.168.1.149:5672/", "rammitmq url")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.GRMAPIService, "v2")
	service.Init(&c.Meta)

	// 初始化消息队列
	mQueue := mq.RabbitMQ{URL: c.MqUrl}
	if err := mQueue.Connect(); err != nil {
		fmt.Print("Faile to connect mq(%s):%s\n", c.MqUrl, err)
		return 1
	}
	defer mQueue.Close()
	c.MsgQueue = &mQueue

	//
	router := mux.NewRouter()
	consulConfig := api.DefaultConfig()
	consulConfig.Address = c.Meta.RegistryAddress
	consulClient, _ := api.NewClient(consulConfig)
	discoveryClient := consul.NewClient(consulClient)

	logger := log.NewLogfmtLogger(os.Stdout)
	ctx := context.Background()
	for svcKey, service := range MakeAPIRouter(ctx) {
		instancer := consul.NewInstancer(discoveryClient, logger, svcKey, []string{}, true)
		for _, method := range service {
			endpointer := sd.NewEndpointer(instancer, method.Factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(3, 30*time.Second, balancer)
			handler := MakeAPIHandler(ctx, retry, logger, svcKey, method.IsPublic)
			router.Methods(method.Method).Path(method.Path).Handler(handler)
		}
	}
	router.Methods("OPTIONS").HandlerFunc(Index)

	// websocket
	hub := newHub()
	go hub.run()

	// 监听任务信息
	c.SubScribeTaskInfo(hub)

	// 任务信息
	router.Methods("GET").Path("/ws/task").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AllowOrigin(w)
		AllowSecurity(w)
		defer r.Body.Close()

		conn, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			logger.Log("method", "ws.task.Upgrade", "error", err.Error())
			return
		}
		defer conn.Close()

		// 解析task_type和task_id
		var taskId string
		taskType := r.URL.Query().Get("task_type")
		if taskType != "" {
			taskId = r.URL.Query().Get("task_id")
		}
		taskType = strings.Replace(taskType, "=", "", -1)
		taskId = strings.Replace(taskId, "=", "", -1)
		fmt.Println("url:", r.URL.String(), ",task_type:", taskType, ",task_id:", taskId)

		ch := make(chan bool)
		client := &Client{hub: hub, conn: conn,
			clientType: "TaskInfo", closed: ch,
			taskType: taskType, taskId: taskId}
		client.hub.register <- client
		defer func() {
			client.hub.unregister <- client
		}()

		if <-ch {
			return
		}
	})

	service.Handle("/", router)
	service.Run()
	return 0
}
