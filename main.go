package main

import (
	"context"
	"github.com/mes1234/immudb/proto/pb"
	"go-micro.dev/v4"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	if ctx.Err() != nil {
		return nil
	}
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("hello world"))

	service.Init()

	err := pb.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		return
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
