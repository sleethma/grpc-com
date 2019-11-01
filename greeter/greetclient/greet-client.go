package main

import (
	"context"
	"fmt"
	"grpc-proto/greeter/greetpb"
	"log"

	"google.golang.org/grpc"
)

type request struct{}

// func (*request) GreetRequest(){

// }

func main() {
	fmt.Println("Client started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreeterClient(cc)
	fmt.Printf("Created client %v \n", c)

	err = unaryCall(c)
	if err != nil {
		log.Panicf("Could not make unary call %v\n", err)
	}
}

func unaryCall(c greetpb.GreeterClient) error {
	req := &greetpb.GreetRequest{
		FirstName: "Donkey",
		LastName:  "Kong",
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		return err
	}
	log.Printf("Response from server %v", res.Result)
	return nil
}
