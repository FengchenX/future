package command

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"

	"data-importer/types"
	"grm-service/log"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Info("socket connect: ", client.clientType)
		case client := <-h.unregister:
			log.Info("socket disconnect:", client.clientType)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
		}
	}
}

//////////////////////////////////////////////////////////////////////////////
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	conn       *websocket.Conn
	clientType string

	closed      chan bool
	retryCount  int
	lastMessage string

	taskType string
	taskId   string
}

func (ws *Client) writeTaskInfo(msg []byte) {
	if ws == nil || ws.lastMessage == string(msg) {
		return
	}

	// 解析task_type,task_id
	var taskInfo types.TaskInfo
	if err := json.Unmarshal(msg, &taskInfo); err != nil {
		log.Error("Failed to parse task info message")
		return
	}

	if len(ws.taskType) > 0 && len(ws.taskId) > 0 {
		if taskInfo.TaskType != ws.taskType || taskInfo.TaskId != ws.taskId {
			return
		}
	} else if len(ws.taskType) > 0 {
		if taskInfo.TaskType != ws.taskType {
			return
		}
	}

	// 发送消息
	if s, ok := ws.hub.clients[ws]; ok && s {
		err := ws.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Warn("Failed to send task info weoskcet:", err)
			ws.retryCount++
			if ws.retryCount >= 3 {
				ws.closed <- true
			}
		}
		ws.lastMessage = string(msg)
	}
}
