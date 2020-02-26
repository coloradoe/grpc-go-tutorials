package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-go-tutorials/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm a Client")

	cc, err := grpc.Dial("Localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephanie",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Roy",
			LastName:  "Maboy",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimesRPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//We've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		fmt.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
