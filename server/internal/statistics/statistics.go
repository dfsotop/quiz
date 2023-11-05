package statistics

import "fmt"

type StatisticsService struct {
	CorrectAnswers   map[int32]string
	historicalScores []float32
}

func NewStatisticsService() *StatisticsService {
	s := &StatisticsService{}
	return s
}

func (s *StatisticsService) CompareWithHistorical(score float32) float32 {
	betterThan := 0
	for _, historicScore := range s.historicalScores {
		if score > historicScore {
			betterThan++
		}
	}

	total := len(s.historicalScores)
	rate := float32(betterThan) / float32(total) * 100
	fmt.Printf("Better than: %d. Total: %d Rate: %.2f%%. Historic scores: %v\n", betterThan, total, rate, s.historicalScores)

	s.historicalScores = append(s.historicalScores, score)

	return rate
}
