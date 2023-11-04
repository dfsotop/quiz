package main

import (
	"log"
	"net"

	pb "github.dfsotop.quiz/m/proto"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuizServiceServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
