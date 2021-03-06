package main

import (
	"flag"
	"fmt"
	pb "github.com/nokamoto/example-ping-service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var (
	port = flag.Int("p", 9000, "grpc server port")
	host = flag.String("h", "localhost", "grpc server host")
)

func main() {
	flag.Parse()

	cc, err := grpc.Dial(fmt.Sprintf("%s:%d", *host, *port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	c := pb.NewPingServiceClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pong, err := c.Send(ctx, &pb.Ping{})
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
}
