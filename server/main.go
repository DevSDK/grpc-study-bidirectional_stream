package main

import (
	"log"
	"net"
	"io"

	"google.golang.org/grpc"
	pb "grpc-example/gen/math"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMathServer
}

func (s *server) Sum(stream pb.Math_SumServer) (error) {
	var sum int32
	sum = 0;
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			log.Print("Client send EOF")
			return nil;
		}
		if err != nil {
			log.Printf("could not recv: %v", err)
			return err
		}
		sum += request.Number;
		log.Printf("Received: %d sum: %d", request.Number, sum);
		if sum % 3 == 0 {
			stream.Send(&pb.Response{Number: sum})
			log.Printf("Send number: %d", sum)
			sum = 0;
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
