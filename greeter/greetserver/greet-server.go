package main

import (
	"context"
	"fmt"
	"grpc-com/greeter/greetpb"
	"log"
	"net"
	"strconv"
	"time"
	"io"
	"google.golang.org/grpc"
)

type servers struct {
}

func (*servers) GreetClientStream(s greetpb.Greeter_GreetClientStreamServer) error{

	fmt.Println("LongGreet Invoked")
	result := "Hello "
	count := 0
	 
	for{
		req, err := s.Recv()
		count++
		if err == io.EOF{
			return s.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}else if err != nil{
			log.Fatalf("error recieving client stream: %v \n", err)
		}
		
		result += "Hello " + req.GetMessageLong() + " " + strconv.Itoa(count) + " \n"
	}
}

func (*servers) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was envoked with %v \n", req)

	firstName := req.GetFirstName()
	result := "Hello " + firstName
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func (*servers) GreetServerStream(req *greetpb.GreetManyTimesRequest, stream greetpb.Greeter_GreetServerStreamServer) error {
	fmt.Printf("GreetServerStream function was envoked with %v \n", req)

	firstNames := req.GetFirstNames()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstNames + " " + strconv.Itoa(i) + " \n"
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	sInstance := grpc.NewServer()
	greetpb.RegisterGreeterServer(sInstance, &servers{})

	if err := sInstance.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
