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

	// doServerStreaming(c)

	doClientStreaming(c)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to do a Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling ComputeAverage RPC: %v", err)
	}

	numbers := []int32{4, 5, 6, 7, 8}
	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receving response from ComputeAverage RPC: %v", err)

	}
	fmt.Printf("the average is: %v\n", res.GetAverage())
}

// create a connection to the server
// grpc.WithInsecure() --> gRPC is SSL secured by default, when we do this we disable SSL security
// grpc.WithInsecure() --> should not be used in production
// cc --> client connection
