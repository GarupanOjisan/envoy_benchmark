package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/garupanojisan/envoy_benchmark/proto/echo"
	"github.com/garupanojisan/envoy_benchmark/server/echo"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	echoServer := &echo.Server{}
	pb.RegisterEchoServerServer(srv, echoServer)

	log.Println("start listening server on :8080")
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
