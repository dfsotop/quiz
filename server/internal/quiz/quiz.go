package quiz

import (
	pb "github.dfsotop.quiz/m/proto"
)

type Service struct {
	QuizTitle      string
	Questions      []*pb.Question
	correctAnswers map[int32]string
}

func NewService(questionsSource Source) *Service {
	quizTitle := questionsSource.GetQuizTitle()
	questions, correctAnswers := questionsSource.GetQuestionsAndAnswers()
	s := &Service{
		Questions:      questions,
		correctAnswers: correctAnswers,
		QuizTitle:      quizTitle,
	}
	return s
}

func (s *Service) GetQuestions() []*pb.Question {
	return s.Questions
}

func (s *Service) GetQuestionsResult(in *pb.RegisterAnswersRequest) []*pb.QuestionResult {
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

func (s *Service) GetScore(answers []*pb.QuestionResult) float32 {
	totalCorrect := 0
	for _, answer := range answers {
		if answer.IsCorrect {
			totalCorrect++
		}
	}
	return float32(totalCorrect) / float32(len(s.Questions))
}
