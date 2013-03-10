package sstats

import (
	"math"
	"testing"
)

func TestMinimum(t *testing.T) {
	tests := []statTest{
		{
			Init:     &Minimum{},
			Values:   nil,
			Valid:    false,
			Expected: 0.0,
		},
		{
			Init:     &Minimum{},
			Values:   []float64{1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Minimum{},
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 0,
		},
		{
			Init:     &Minimum{},
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 0,
		},
		{
			Init:     &Minimum{},
			Values:   []float64{math.Inf(-1), -math.MaxFloat64},
			Valid:    true,
			Expected: math.Inf(-1),
		},
	}

	runTests(t, tests)
}
