package main

import (
	"io"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"github.com/feng/future/grpc/protocol"
	"golang.org/x/net/context"
)
func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("创建grpc conn失败**********", err)	
	}
	defer conn.Close()
	client := protocol.NewRouteGuideClient(conn)
	_, err = client.GetFeature(context.Background(), &protocol.Point{Latitude: 10, Longitude: 20})
	if err != nil {
		log.Fatalln(err)
	}
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
	}
}
