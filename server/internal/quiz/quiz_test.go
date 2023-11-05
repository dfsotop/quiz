package quiz

import (
	"testing"

	pb "github.dfsotop.quiz/m/proto"
)

func TestGetQuestions(t *testing.T) {
	// Given
	source := newMockSource()
	expectedQuestions, _ := source.GetQuestionsAndAnswers()

	service := NewQuizService(source)

	// When
	actualQuestions := service.GetQuestions()

	// Then
	if len(expectedQuestions) != len(actualQuestions) {
		t.Errorf("expected %d questions, got %d", len(expectedQuestions), len(actualQuestions))
	}

	for i, question := range expectedQuestions {
		if question != actualQuestions[i] {
			t.Errorf("expected question %d to be %v, got %v", i, question, actualQuestions[i])
		}
	}
}

func TestQuizAnswers(t *testing.T) {
	// Given
	source := newMockSource()
	_, answers := source.GetQuestionsAndAnswers()
	service := NewQuizService(source)

	if len(answers) != len(service.correctAnswers) {
		t.Errorf("expected %d answers, got %d", len(answers), len(service.Questions))
	}
}

func TestAllAnswersAreInCorrect(t *testing.T) {
	source := newMockSource()
	service := NewQuizService(source)

	wrongAnswers := make(map[int32]string)
	for _, question := range service.Questions {
		wrongAnswers[question.Id] = "c"
	}

	answers := make([]*pb.Answer, 0, len(wrongAnswers))
	for questionId, answer := range wrongAnswers {
		answers = append(answers, &pb.Answer{
			QuestionId:     questionId,
			ChosenOptionId: answer,
		})
	}

	req := &pb.RegisterAnswersRequest{
		Answers: answers,
	}

	result := service.GetQuestionsResult(req)

	for _, questionResult := range result {
		if questionResult.IsCorrect {
			t.Errorf("expected question %d to be incorrect, got correct", questionResult.QuestionId)
		}
	}
}

func TestAllAnswersAreCorrect(t *testing.T) {
	source := newMockSource()
	service := NewQuizService(source)

	correctAnswers := make(map[int32]string)
	for _, question := range service.Questions {
		correctAnswers[question.Id] = service.correctAnswers[question.Id]
	}

	answers := make([]*pb.Answer, 0, len(correctAnswers))
	for questionId, answer := range correctAnswers {
		answers = append(answers, &pb.Answer{
			QuestionId:     questionId,
			ChosenOptionId: answer,
		})
	}

	req := &pb.RegisterAnswersRequest{
		Answers: answers,
	}

	result := service.GetQuestionsResult(req)

	for _, questionResult := range result {
		if !questionResult.IsCorrect {
			t.Errorf("expected question %d to be incorrect, got correct", questionResult.QuestionId)
		}
	}
}

type mockSource struct {
	Questions      []*pb.Question
	CorrectAnswers map[int32]string
}

func newMockSource() *mockSource {
	s := &mockSource{}
	s.Questions = []*pb.Question{
		{
			Id:   1,
			Text: "Question 1",
			Options: []*pb.Option{
				{Id: "a", Text: "Option 1"},
				{Id: "b", Text: "Option 2"},
			},
		},
		{
			Id:   2,
			Text: "Question 2",
			Options: []*pb.Option{
				{Id: "a", Text: "Option 1"},
				{Id: "b", Text: "Option 2"},
			},
		},
	}
	s.CorrectAnswers = map[int32]string{
		1: "a",
		2: "b",
	}
	return s
}

func (s *mockSource) GetQuestionsAndAnswers() ([]*pb.Question, map[int32]string) {
	return s.Questions, s.CorrectAnswers
}
