package sstats

import (
	"math"
	"testing"
)

func TestMaximum(t *testing.T) {
	tests := []statTest{
		{
			Init:     &Maximum{},
			Values:   nil,
			Valid:    false,
			Expected: 0.0,
		},
		{
			Init:     &Maximum{},
			Values:   []float64{1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Maximum{},
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 5,
		},
		{
			Init:     &Maximum{},
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 2,
		},
		{
			Init:     &Maximum{},
			Values:   []float64{math.Inf(-1), -math.MaxFloat64},
			Valid:    true,
			Expected: -math.MaxFloat64,
		},
	}

	runTests(t, tests)
}
