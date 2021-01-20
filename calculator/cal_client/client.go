package main

import (
	"context"
	fmt "fmt"
	"go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
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
	doUnary(c)
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
