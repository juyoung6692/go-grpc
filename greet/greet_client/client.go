package main

import (
	"context"
	"fmt"
	"go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main(){
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err!=nil{
		log.Fatalf("could not connet: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient)  {
	fmt.Printf("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "JuYoung",
			LastName: "Lee",
		},
	}
	res, err :=c.Greet(context.Background(), req)
	if err!=nil{
		log.Fatalf("erro while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v ", res.Result)

}
