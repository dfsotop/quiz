package quiz

import pb "github.dfsotop.quiz/m/proto"

type Source interface {
	GetQuestionsAndAnswers() ([]*pb.Question, map[int32]string)
}

type WorldCupSource struct {
	Questions      []*pb.Question
	CorrectAnswers map[int32]string
}

func NewWorldCupSource() *WorldCupSource {
	s := &WorldCupSource{}
	s.createQuestionsAndAnswers()
	return s
}

func (s *WorldCupSource) createQuestionsAndAnswers() {
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
}

func (s *WorldCupSource) GetQuestionsAndAnswers() ([]*pb.Question, map[int32]string) {
	return s.Questions, s.CorrectAnswers
}
