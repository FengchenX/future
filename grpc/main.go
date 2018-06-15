package main

import (
	"fmt"
	"github.com/feng/future/grpc/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

func main() {
	grpcSvr := grpc.NewServer()
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	mySvr := mySvr{}
	protocol.RegisterRouteGuideServer(grpcSvr, &mySvr)
	if err = grpcSvr.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

type mySvr struct{}

var feats = []protocol.Feature{
	{Name: "feng", Location: &protocol.Point{Latitude: 10, Longitude: 20}},
	{Name: "chen", Location: &protocol.Point{Latitude: 30, Longitude: 80}},
}

//简单rpc
func (svr *mySvr) GetFeature(ctx context.Context, point *protocol.Point) (*protocol.Feature, error) {
	fmt.Println(point.Latitude, point.Longitude)
	return &protocol.Feature{}, nil
}

//服务端流式rpc
func (svr *mySvr) ListFeatures(rect *protocol.Rectangle, stream protocol.RouteGuide_ListFeaturesServer) error {
	fmt.Println("list******", *rect)
	for _, feat := range feats {
		stream.Send(&feat)
	}
	return nil
}

//客户端流式rpc
func (svr *mySvr) RecordRoute(stream protocol.RouteGuide_RecordRouteServer) error {
	fmt.Println("客户端流")
	for {
		p, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("svr******", *p)
			return stream.SendAndClose(&protocol.RouteSummary{})
		}
		return nil
	}
}

//双向流式rpc
func (svr *mySvr) RouteChat(protocol.RouteGuide_RouteChatServer) error {
	return nil
}
