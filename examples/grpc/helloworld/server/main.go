package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dayu-go/gkit/app"
	"github.com/dayu-go/gkit/examples/grpc/helloworld/pb"
	"github.com/dayu-go/gkit/transport/grpc"
)

var (
	Name    = "helloworld"
	Version = "v1.0.0"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "error" {
		return nil, fmt.Errorf("invalid argument %s", in.Name)
	}
	if in.Name == "panic" {
		panic("server panic")
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func main() {
	s := &server{}
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
	)
	pb.RegisterGreeterServer(grpcSrv, s)

	app := app.New(
		app.Name(Name),
		app.Server(
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

}
