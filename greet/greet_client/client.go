package main

import (
	"fmt"
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
	fmt.Printf("Created client: %f", c)

}
