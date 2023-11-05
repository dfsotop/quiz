package statistics

import (
	"fmt"
	"sync"
)

type Service struct {
	historicalScores []float32
	mu               sync.Mutex
}

func NewService() *Service {
	s := &Service{}
	return s
}

func (s *Service) CompareWithHistorical(score float32) float32 {
	s.mu.Lock()
	defer s.mu.Unlock()

	betterThan := 0
	for _, historicScore := range s.historicalScores {
		if score > historicScore {
			betterThan++
		}
	}

	total := len(s.historicalScores)
	var rate float32 = 100
	if total != 0 {
		rate = float32(betterThan) / float32(total) * 100
	}
	fmt.Printf("Better than: %d. Total: %d Rate: %.2f%%. Historical scores: %v\n", betterThan, total, rate, s.historicalScores)

	s.historicalScores = append(s.historicalScores, score)

	return rate
}
