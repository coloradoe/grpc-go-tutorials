package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

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

	var k int64 = 2
	N := req.GetNp().GetX()
	for N > 1 {
		if (N % k) == 0 {
			fmt.Println(k)
			N = N / k
			result := N
			res := &calculatorpb.PrimeResponse{
				Result: result,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k = k + 1
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
