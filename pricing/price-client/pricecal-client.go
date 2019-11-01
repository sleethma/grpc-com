package main

import (
	"context"
	"fmt"
	"grpc-proto/pricing/priceProtos"
	"log"

	"google.golang.org/grpc"
)

type request struct {
	numA int32
	numB int32
}

func main() {
	fmt.Println("Starting client")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failure getting client connection %v \n", err)
	}
	defer conn.Close()

	c := priceProtos.NewPricerClient(conn)

	req := &priceProtos.PriceRequest{
		PriceItemA: 5,
		PriceItemB: 10,
	}

	res, err := c.GetPrice(context.Background(), req)

	fmt.Printf("Response to client %d\n", res)

}
