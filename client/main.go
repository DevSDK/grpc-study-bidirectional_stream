package main

import (
	"context"
	"log"
	"time"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	pb "grpc-example/gen/math"
)

const (
	address     = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewMathClient(conn)

	// Sum Deadline을 5분으로 설정합니다. 5분간 유효합니다. (6장 참조)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute * 5)
	
	defer conn.Close()
	defer cancel()
	
	// GRPC CALL
	stream, err := c.Sum(ctx)
	
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}
	log.Print("Client is successfully connected")

	var input int32
	channel := make(chan struct{})
	go receiver(stream, channel)
	for {
		fmt.Print("Number (exit -1): ") 
		fmt.Scanf("%d", &input)

		if input == -1{
			if err := stream.CloseSend(); err != nil {
				log.Fatal(err)
			}
			break;
		}
		sendNumber(input, stream)
	}
	<- channel
}

func sendNumber(number int32, stream pb.Math_SumClient) {
	if err := stream.Send(&pb.Request{Number:number}); err != nil {
		log.Printf("Could not send: %v", err)
	}
}

func receiver(stream pb.Math_SumClient, channel chan struct{}) {
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Print("Server send EOF")
			break;
		}
		if err != nil {
			if status.Code(err) == codes.DeadlineExceeded {
				log.Printf("Deadline is exceeded")
				break;
			}
			log.Printf("receiver error : %v", err)
		}
		log.Printf("Received: %d", response.Number)
	}
	close(channel)
}