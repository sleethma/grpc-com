package main

import (
	"context"
	"fmt"
	"grpc-proto/pricing/priceProtos"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GetPrice(ctx context.Context, req *priceProtos.PriceRequest) (*priceProtos.PriceResponse, error) {
	priceA := req.GetPriceItemA()
	priceB := req.GetPriceItemB()
	totalPrice := priceA + priceB

	res := priceProtos.PriceResponse{
		TotalPrice: totalPrice,
	}
	return &res, nil
}

func main() {
	fmt.Println("Starting server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to start listener %v \n", err)
	}

	s := grpc.NewServer()
	priceProtos.RegisterPricerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server not serving %v \n", err)
	}
}
