package main
import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/peteratkinson/learn-grpc/hello"
)

func main() {
	var conn  *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial gRPC connection %s", err)
	}

	// defer the close connection until the service terminates
	defer conn.Close()

	// create a new client that will send procedural calls to the gRPC service
	helloServiceClient := hello.NewHelloServiceClient(conn)
	
	request := hello.HelloRequest{ Name: "Peter" }

	// call the SayHello procedural call
	response , err := helloServiceClient.SayHello(context.Background(),&request)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	response, err = helloServiceClient.SayHelloAgain(context.Background(), &request)

	if err != nil {
		log.Fatalf("Error when calling SayHelloAgain: %s", err)
	}

	log.Printf("Response from Server: %s", response.Message)
}