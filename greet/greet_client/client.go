package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MelvinKim/go-gRPC-intro/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am a client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client could not connect to the server: %v", err)
	}

	defer cc.Close()

	// create a client with the connection above
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a Unary RPC")

	// create a greet request
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stepahen",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", res)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

// create a connection to the server
// grpc.WithInsecure() --> gRPC is SSL secured by default, when we do this we disable SSL security
// grpc.WithInsecure() --> should not be used in production
// cc --> client connection
