package main

import (
	"context"
	"fmt"
	"go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello" + firstName + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	result := "Hello "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + ", "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone  function was invoked with a streaming request\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello" + firstName + "!"

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Faild to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
