package quiz

import pb "github.dfsotop.quiz/m/proto"

type Source interface {
	GetQuizTitle() string
	GetQuestionsAndAnswers() ([]*pb.Question, map[int32]string)
}

type WorldCupSource struct {
	title          string
	Questions      []*pb.Question
	CorrectAnswers map[int32]string
}

func NewWorldCupSource() *WorldCupSource {
	s := &WorldCupSource{
		title: "Last champions of the FIFA World Cup",
	}
	s.createQuestionsAndAnswers()
	return s
}

func (s *WorldCupSource) GetQuizTitle() string {
	return s.title
}

func (s *WorldCupSource) createQuestionsAndAnswers() {
	s.Questions = []*pb.Question{
		{
			Id:   1,
			Text: "Who won the 2022 FIFA World Cup?",
			Options: []*pb.Option{
				{Id: "a", Text: "Argentina"},
				{Id: "b", Text: "Croatia"},
				{Id: "c", Text: "Germany"},
				{Id: "d", Text: "Brazil"},
			},
		},
		{
			Id:   2,
			Text: "Who won the 2018 FIFA World Cup?",
			Options: []*pb.Option{
				{Id: "a", Text: "France"},
				{Id: "b", Text: "Spain"},
				{Id: "c", Text: "Italy"},
				{Id: "d", Text: "Spain"},
			},
		},
		{
			Id:   2,
			Text: "Who won the 2014 FIFA World Cup?",
			Options: []*pb.Option{
				{Id: "a", Text: "Germany"},
				{Id: "b", Text: "Spain"},
				{Id: "c", Text: "Italy"},
				{Id: "d", Text: "Brazil"},
			},
		},
	}
	s.CorrectAnswers = map[int32]string{
		1: "a",
		2: "a",
		3: "a",
	}
}

func (s *WorldCupSource) GetQuestionsAndAnswers() ([]*pb.Question, map[int32]string) {
	return s.Questions, s.CorrectAnswers
}
