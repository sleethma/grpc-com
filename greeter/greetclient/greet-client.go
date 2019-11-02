package main

import (
	"context"
	"fmt"
	"grpc-com/greeter/greetpb"
	"io"
	"log"
	"google.golang.org/grpc"
)

type request struct{}

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

	// Server streaming
	// err = serverStreamCall(c)
	// if err != nil {
	// 	log.Panicf("Could not make unary call %v\n", err)
	// }

	// Client streaming
	err2 := streamClientReq(c)
	if err2 != nil{
		log.Fatalf("Client failed with %v \n", err)
	}
}

func streamClientReq(c greetpb.GreeterClient) error{
	firstName := "Starter" 

	stream, err := c.GreetClientStream(context.Background())
	if err != nil{
		return err
		}

	for i := 0; i < 10; i++{
		firstName += "r"
	req := &greetpb.LongGreetRequest{
		MessageLong : firstName,
	}

	stream.Send(req)
	// time.Sleep(3 * time.Second)
	}

	resp , err := stream.CloseAndRecv()
	if err != nil{
		return err
	}
	fmt.Println(resp.Result)

return nil
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
