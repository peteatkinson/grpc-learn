package hello


import (
	"fmt"
	"log"
	"golang.org/x/net/context"
)

// Server that will implement the HelloServiceServer interface
type Server struct {
	UnimplementedHelloServiceServer
}

// implementation for each procedural call the gRPC service contains
func (s *Server) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	log.Printf("Received message body from client: %s", request.Name)
	return &HelloReply{ Message: fmt.Sprintf("Hello %s from the gRPC service!", request.Name)}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, request* HelloRequest) (*HelloReply, error) {
	log.Printf("Received message body from client: %s", request.Name)
	return &HelloReply{Message: fmt.Sprintf("Hello again %s from the gRPC service!", request.Name)}, nil
}