package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MelvinKim/go-gRPC-intro/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("calculator client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client could not connect to the server: %v", err)
	}

	defer cc.Close()

	// create a client with the connection above
	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to do a Sum Unary RPC...")

	// create a greet request
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", res)
	}
	log.Printf("Response from Greet: %v", res.SumResult)
}

// create a connection to the server
// grpc.WithInsecure() --> gRPC is SSL secured by default, when we do this we disable SSL security
// grpc.WithInsecure() --> should not be used in production
// cc --> client connection
