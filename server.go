package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/peteratkinson/learn-grpc/hello"
)

func main() {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	// create a new gRPC server
	grpcServer := grpc.NewServer()

	// create the new gRPC server implementation
	helloServer := hello.Server{}

	// then register the gPRC service to the server
	hello.RegisterHelloServiceServer(grpcServer, &helloServer)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}
}