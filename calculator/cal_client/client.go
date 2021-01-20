package main

import (
	"context"
	fmt "fmt"
	"go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connet: #{err}")
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	fmt.Printf("Created client: %v", c)
	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a Unary RPC...\n")
	var num1, num2 int32
	fmt.Scanf("%d,%d", &num1, &num2)
	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			Number1: num1,
			Number2: num2,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}

	log.Printf("number1 + number2 =  %v ", res.Result)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starginf to do Server a Server Streaming RPC...\n")
	var number int32
	fmt.Scanf("%d", &number)
	req := &calculatorpb.NumberRequest{
		Number: number,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Streaming Server RPC: %V", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while calling Server Streaming RPC: %v", err)
		}
		log.Printf("PrimeNumberDecomposition(%d): %v", number, msg.GetResult())
	}

}
