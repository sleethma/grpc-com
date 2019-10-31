package main

import (
	"fmt"
	"grpc-proto/greeter/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreeterClient(cc)
	fmt.Printf("Created client %v \n", c)
}
