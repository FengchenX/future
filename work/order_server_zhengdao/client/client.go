package client

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"reflect"
	"sub_account_service/order_server_zhengdao/lib"
	"sub_account_service/order_server_zhengdao/protocol"
)

var Cli *Client

// grpc Client
type Client struct {
	C     protocol.ApiServiceClient                // grpc client
	Conn  chan proto.Message                       // msg chan
	Addr  string                                   // client addr
	funcs map[reflect.Type]func(msg proto.Message) // functionmap
}

// set-up new client for grpc connection
func NewClient(addr string) *Client {
	glog.Infoln(lib.Log("client", "", "NewClient"), "starting client addr:", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		glog.Errorln(lib.Log("client err", "", "NewClient"), "Can't connect: "+addr, err)
	}
	client := &Client{
		Addr:  addr,
		Conn:  make(chan proto.Message, 500),
		funcs: make(map[reflect.Type]func(msg proto.Message), 0),
	}
	client.C = protocol.NewApiServiceClient(conn)
	Cli = client
	return client
}

// register func for call back
func (this *Client) Register(f func(proto.Message), msg proto.Message) {
	tpy := reflect.TypeOf(msg)
	if _, exist := this.funcs[tpy]; !exist {
		this.funcs[tpy] = f
	}
}

// read loop&event loop
func (this *Client) Run() {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln("client run err ", err)
		}
	}()
	for msg := range this.Conn {
		go func(m proto.Message) {
			tpy := reflect.TypeOf(m)
			if f, exist := this.funcs[tpy]; exist {
				f(m)
			}
		}(msg)
	}
}
