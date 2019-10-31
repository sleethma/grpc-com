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
	firstName := req.GetFirstName()
	result := "Hello " + firstName
	res
	return 
}

Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
XXX_NoUnkeyedLiteral struct{} `json:"-"`
XXX_unrecognized     []byte   `json:"-"`
XXX_sizecache        int32    `json:"-"`

func main() {
	fmt.Println("starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
