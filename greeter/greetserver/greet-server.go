package main

import (
	"context"
	"fmt"
	"grpc-proto/greeter/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was envoked with %v \n", req)

	firstName := req.GetFirstName()
	result := "Hello " + firstName
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func main() {
	fmt.Println("starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	sInstance := grpc.NewServer()
	greetpb.RegisterGreeterServer(sInstance, &server{})

	if err := sInstance.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
