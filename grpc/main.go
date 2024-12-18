package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"example.com/test/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	go startServer()
	time.Sleep(2 * time.Second)
	go startClient()
	time.Sleep(2 * time.Second)
}

func startServer() {
	helloServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(helloServer, new(proto.HelloServiceServerImp))
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	helloServer.Serve(listen)
}

func startClient() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewHelloServiceClient(conn)
	reply, err := client.SayHello(context.Background(), &proto.String{Value: "李白"})
	if err != nil {
		panic(err)
	}
	fmt.Println("oytp reply = ", reply.GetValue())
}
