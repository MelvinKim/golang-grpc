package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MelvinKim/go-gRPC-intro/calculator/calculatorpb"
	"github.com/MelvinKim/go-gRPC-intro/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.GreetServiceServer
	calculatorpb.CalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("calculator server")

	// create a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
