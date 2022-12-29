package main

import (
	"context"
	"fmt"
	"io"
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

	// doUnary(c)

	doServerStreaming(c)
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
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v\n", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to do a PrimeDecomposition Server Streaming RPC...")

	// create a greet request
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 1247473939238,
	}
	strean, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v\n", err)
	}
	for {
		res, err := strean.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened in PrimeDecomposition RPC: %v\n", err)
		}
		fmt.Printf("server streaming response Prime Factor: %v\n", res.GetPrimeFactor())
	}
}

// create a connection to the server
// grpc.WithInsecure() --> gRPC is SSL secured by default, when we do this we disable SSL security
// grpc.WithInsecure() --> should not be used in production
// cc --> client connection
