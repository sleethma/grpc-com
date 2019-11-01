package main

import (
	"context"
	"fmt"
	"grpc-proto/greeter/greetpb"
	"io"
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

	// err = unaryCall(c)
	// if err != nil {
	// 	log.Panicf("Could not make unary call %v\n", err)
	// }

	err = serverStreamCall(c)
	if err != nil {
		log.Panicf("Could not make unary call %v\n", err)
	}
}

func serverStreamCall(c greetpb.GreeterClient) error {
	req := &greetpb.GreetManyTimesRequest{
		FirstNames: "Multiman",
		LastNames:  "ExpectingMore",
	}
	resp, err := c.GreetServerStream(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Printf("Multiple response: %v \n", msg.GetResult())
	}
	return nil
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
