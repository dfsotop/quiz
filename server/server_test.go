package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.dfsotop.quiz/m/proto"
	"github.dfsotop.quiz/m/server/internal/quiz"
)

func TestRegisterAnswers_InvalidNumberOfAnswers(t *testing.T) {
	// Given
	s := NewServer(quiz.NewWorldCupSource())

	answers := make([]*pb.Answer, 0, 1)

	answers = append(answers, &pb.Answer{
		QuestionId:     1,
		ChosenOptionId: "a",
	})

	req := &pb.RegisterAnswersRequest{
		Answers: answers,
	}
	_, err := s.RegisterAnswers(context.Background(), req)

	assert.Error(t, err)
}

func TestRegisterAnswers_SuccessfulProcessing(t *testing.T) {
	// Given
	s := NewServer(quiz.NewWorldCupSource())

	myAnswers := make(map[int32]string)
	for _, question := range s.quizService.GetQuestions() {
		myAnswers[question.Id] = "c"
	}

	answers := make([]*pb.Answer, 0, len(myAnswers))
	for questionId, answer := range myAnswers {
		answers = append(answers, &pb.Answer{
			QuestionId:     questionId,
			ChosenOptionId: answer,
		})
	}

	req := &pb.RegisterAnswersRequest{
		Answers: answers,
	}

	// When
	resp, err := s.RegisterAnswers(context.Background(), req)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, (float32)(0.0), resp.Score)
	assert.Equal(t, len(s.quizService.Questions), len(resp.QuesitonsResults))
	assert.Equal(t, float32(100), resp.Statistics.BetterThan)
}
