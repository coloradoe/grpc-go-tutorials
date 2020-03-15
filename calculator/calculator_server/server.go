package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-go-tutorials/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Printf("Received Sum RPC: %v", req)

	a := req.GetSum().GetA()
	b := req.GetSum().GetB()
	sum := a + b

	res := &calculatorpb.SumResponse{
		Result: sum,
	}

	return res, nil

}

func (s *server) NumberPrime(req *calculatorpb.PrimeRequest, stream calculatorpb.SumService_NumberPrimeServer) error {

	fmt.Printf("Received Number RPC: %v", req)

	number := req.GetNp().GetX()
	divisor := int64(2)

	for number > 1 {
		if (number % divisor) == 0 {
			fmt.Println(divisor)
			stream.Send(&calculatorpb.PrimeResponse{
				Result: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to: %v", divisor)
		}
	}
	return nil

}

func main() {
	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to liste: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
