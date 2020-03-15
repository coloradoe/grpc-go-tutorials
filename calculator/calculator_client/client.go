package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-go-tutorials/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I'm a Client")

	cc, err := grpc.Dial("Localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	// c := calculatorpb.NewSumServiceClient(cc)
	c := calculatorpb.NewSumServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	doServerStreaming(c)

}

func doUnary(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			A: 3,
			B: 6,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Result)
}

func doServerStreaming(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do a Server Stream RPC...")

	req := &calculatorpb.PrimeRequest{
		Np: 12,
	}
	resStream, err := c.NumberPrime(context.Background(), req)
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
		fmt.Printf("Response from NumberPrime: %v", msg.GetResult())
	}
}
