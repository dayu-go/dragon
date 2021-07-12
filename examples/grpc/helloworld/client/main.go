package main

import (
	"context"
	"log"

	"github.com/dayu-go/gkit/examples/grpc/helloworld/pb"
	transgrpc "github.com/dayu-go/gkit/transport/grpc"
)

func main() {
	callGRPC()
}

func callGRPC() {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:9000"),
	)
	if err != nil {
		panic(err)
	}
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)

	// returns error
	_, err = client.SayHello(context.Background(), &pb.HelloRequest{Name: "error"})
	if err != nil {
		log.Printf("[grpc] SayHello error: %v\n", err)
	}
}
