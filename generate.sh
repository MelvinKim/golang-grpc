#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=. greet/greetpb/greet.proto
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=. calculator/calculatorpb/calculator.proto
# the command is responsible for generating the protobuf definitions and also generating gRPC code



# magic commands
# protoc --go-grpc_out=. calculator/calculatorpb/calculator.proto
# protoc --go_out=. --go-grpc_out=. calculator/calculatorpb/calculator.proto