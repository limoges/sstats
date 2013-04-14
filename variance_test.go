package sstats

import "testing"

func TestVariance(t *testing.T) {
	tests := []statTest{
		{
			Init:     &Variance{},
			Values:   nil,
			Valid:    false,
			Expected: 0.0,
		},
		{
			Init:     &Variance{},
			Values:   []float64{1},
			Valid:    true,
			Expected: 0,
		},
		{
			Init:     &Variance{},
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 3.5,
		},
		{
			Init:     &Variance{},
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Variance{},
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 2,
		},
	}

	runTests(t, tests)
}
