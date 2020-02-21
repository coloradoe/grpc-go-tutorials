package main

import (
	"context"
	"fmt"
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

	c := calculatorpb.NewSumServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	doUnary(c)

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
