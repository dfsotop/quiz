package quiz

import (
	pb "github.dfsotop.quiz/m/proto"
)

type QuizService struct {
	Questions      []*pb.Question
	correctAnswers map[int32]string
}

func NewQuizService() *QuizService {
	s := &QuizService{}
	s.createQuestionsAndAnswers()
	return s
}

func (s *QuizService) createQuestionsAndAnswers() {
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
	s.correctAnswers = map[int32]string{
		1: "a",
		2: "b",
	}
}

func (s *QuizService) GetQuestions() []*pb.Question {
	return s.Questions
}

func (s *QuizService) GetQuestionsResult(in *pb.RegisterAnswersRequest) []*pb.QuestionResult {
	questionResults := make([]*pb.QuestionResult, 0)
	for _, answer := range in.Answers {
		questionResult := &pb.QuestionResult{
			QuestionId:      answer.QuestionId,
			ChosenOptionId:  answer.ChosenOptionId,
			CorrectOptionId: s.correctAnswers[answer.QuestionId],
			IsCorrect:       s.correctAnswers[answer.QuestionId] == answer.ChosenOptionId,
		}
		questionResults = append(questionResults, questionResult)
	}
	return questionResults
}

func (s *QuizService) GetScore(answers []*pb.QuestionResult) float32 {
	totalCorrect := 0
	for _, answer := range answers {
		if answer.IsCorrect {
			totalCorrect++
		}
	}
	return float32(totalCorrect) / float32(len(s.Questions))
}
