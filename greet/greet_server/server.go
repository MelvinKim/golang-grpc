package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MelvinKim/go-gRPC-intro/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.GreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)

	// extracting information from our request
	firstName := req.GetGreeting().GetFirstName()

	// create the response
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello world")

	// create a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
