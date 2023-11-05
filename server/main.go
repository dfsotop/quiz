package main

import (
	"log"
	"net"

	pb "github.dfsotop.quiz/m/proto"
	"github.dfsotop.quiz/m/server/internal/quiz"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuizServiceServer(s, newServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newServer() *server {
	questionsSource := quiz.NewWorldCupSource()

	srv := NewServer(questionsSource)
	return srv
}
