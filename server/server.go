package main

import (
	"context"
	"fmt"

	pb "github.dfsotop.quiz/m/proto"
)

type server struct {
	pb.UnimplementedQuizServiceServer
	questions      []*pb.Question
	correctAnswers map[int32]int32
}

func NewServer() *server {
	s := &server{}
	s.createQuestions()
	return s
}

func (s *server) createQuestions() {
	s.questions = []*pb.Question{
		{
			Id:   1,
			Text: "Question 1",
			Options: []*pb.Option{
				{Id: 1, Text: "Option 1"},
				{Id: 2, Text: "Option 2"},
			},
		},
		{
			Id:   2,
			Text: "Question 2",
			Options: []*pb.Option{
				{Id: 1, Text: "Option 1"},
				{Id: 2, Text: "Option 2"},
			},
		},
	}
	s.correctAnswers = map[int32]int32{
		1: 1,
		2: 2,
	}
}

func (s *server) GetQuestions(ctx context.Context, in *pb.GetQuestionsRequest) (*pb.GetQuestionsResponse, error) {
	return &pb.GetQuestionsResponse{Questions: s.questions}, nil
}

func (s *server) RegisterAnswers(ctx context.Context, in *pb.RegisterAnswersRequest) (*pb.RegisterAnswersResponse, error) {
	for _, answer := range in.Answers {
		fmt.Printf(answer.String())
	}
	return &pb.RegisterAnswersResponse{QuizRating: 12}, nil
}
