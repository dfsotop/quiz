package main

import (
	"context"
	"fmt"

	pb "github.dfsotop.quiz/m/proto"
	"github.dfsotop.quiz/m/server/internal/quiz"
	"github.dfsotop.quiz/m/server/internal/statistics"
)

type server struct {
	pb.UnimplementedQuizServiceServer

	quizService       *quiz.QuizService
	statisticsService *statistics.StatisticsService
}

func NewServer(questionsSource quiz.Source) *server {
	s := &server{}
	s.quizService = quiz.NewQuizService(questionsSource)
	s.statisticsService = statistics.NewStatisticsService()
	return s
}

func (s *server) GetQuestions(ctx context.Context, in *pb.GetQuestionsRequest) (*pb.GetQuestionsResponse, error) {
	questions := s.quizService.GetQuestions()
	return &pb.GetQuestionsResponse{Questions: questions}, nil
}

func (s *server) RegisterAnswers(ctx context.Context, in *pb.RegisterAnswersRequest) (*pb.RegisterAnswersResponse, error) {
	fmt.Printf("Registering answers:\n")
	for _, answer := range in.Answers {
		fmt.Printf(answer.String())
	}
	fmt.Printf("\n")

	if len(in.Answers) != len(s.quizService.Questions) {
		return nil, fmt.Errorf("invalid number of answers")
	}

	questionsResuls := s.quizService.GetQuestionsResult(in)
	score := s.quizService.GetScore(questionsResuls)
	betterThan := s.statisticsService.CompareWithHistorical(score)

	statistics := &pb.Statistics{
		BetterThan: float32(betterThan),
	}
	return &pb.RegisterAnswersResponse{
		Score:            score,
		QuesitonsResults: questionsResuls,
		Statistics:       statistics,
	}, nil

}
