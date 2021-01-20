package main

import (
	"context"
	"fmt"
	"go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Caculator function was invoked with: %v", req)
	number1 := req.GetCalculator().Number1
	number2 := req.GetCalculator().Number2
	result := number1 + number2
	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.NumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	var p_number int32
	p_number = 2
	number := req.GetNumber()
	for number > 1 {
		if number%p_number == 0 {
			number = number / p_number
			result := p_number
			res := &calculatorpb.PrimeNumberResponse{
				Result: result,
			}
			stream.Send(res)
			time.Sleep(500 * time.Millisecond)

		} else {
			p_number++
		}
	}

	return nil
}

func main() {
	fmt.Println("This is calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: #{err}")
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: #{err}")
	}
}
