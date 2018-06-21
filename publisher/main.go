package main

import (
	"context"
	"log"
	"net"

	pb "github.com/faisalburhanudin/pubsub/ping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8000"
)

type server struct{}

func (s *server) PingMe(ctx context.Context, in *pb.PingRequest, stream pb.) error {
	return 
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
