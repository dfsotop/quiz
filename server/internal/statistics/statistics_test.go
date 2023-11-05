package statistics

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareWithHistorical_LowerThanAll(t *testing.T) {
	// Given
	historicalScores := []float32{10, 20, 30, 40, 50}
	service := &Service{
		historicalScores: historicalScores,
	}

	// When
	rate := service.CompareWithHistorical(5)

	// Then
	expectedRate := float32(0)
	assert.Equal(t, expectedRate, rate)
	assert.Equal(t, []float32{10, 20, 30, 40, 50, 5}, service.historicalScores)
}

func TestCompareWithHistorical_HigherThanAll(t *testing.T) {
	// Given
	historicalScores := []float32{10, 20, 30, 40, 50}
	service := &Service{
		historicalScores: historicalScores,
	}

	// When
	rate := service.CompareWithHistorical(55)

	// Then
	expectedRate := float32(100)
	assert.Equal(t, expectedRate, rate)
	assert.Equal(t, []float32{10, 20, 30, 40, 50, 55}, service.historicalScores)
}

func TestCompareWithHistorical_NoHistoricalScores(t *testing.T) {
	// Given
	service := NewService()

	// When
	rate := service.CompareWithHistorical(55)

	// Then
	expectedRate := float32(100)
	assert.Equal(t, expectedRate, rate)
	assert.Equal(t, []float32{55}, service.historicalScores)
}

func TestCompareWithHistorical_Concurrent(t *testing.T) {
	//Given
	s := NewService()
	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// When
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			s.CompareWithHistorical(float32(i))
		}()
	}

	// Then
	wg.Wait()

	if len(s.historicalScores) != numGoroutines {
		t.Errorf("expected %d historical scores, got %d", numGoroutines, len(s.historicalScores))
	}
}
