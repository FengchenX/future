package main

import (
	//"io"
	//"fmt"
	"github.com/feng/future/grpc/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("创建grpc conn失败**********", err)
	}
	defer conn.Close()
	client := protocol.NewRouteGuideClient(conn)

	//简单流式rpc
	_, err = client.GetFeature(context.Background(), &protocol.Point{Latitude: 10, Longitude: 20})
	if err != nil {
		log.Fatalln(err)
	}

	/*
		//服务端流式rpc
		feats, err := client.ListFeatures(context.Background(),
		&protocol.Rectangle{Lo: &protocol.Point{Latitude: 10, Longitude: 20}, Hi: &protocol.Point{Latitude: 20, Longitude: 40}})
		if err != nil {
			log.Fatalln(err)
		}
		for {
			feat, err := feats.Recv()
			if err == io.EOF {
				return
			}
			fmt.Println("cli******", *feat)
		}*/

	//客户端流式rpc
	rs, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 5; i++ {
		err = rs.Send(&protocol.Point{Latitude: 10, Longitude: 100})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
