package sstats

import "testing"

func TestMedian(t *testing.T) {
	tests := []statTest{
		{
			Init:     &Median{},
			Values:   nil,
			Valid:    false,
			Expected: 0.0,
		},
		{
			Init:     &Median{},
			Values:   []float64{1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Median{},
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 3,
		},
		{
			Init:     &Median{},
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Median{},
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 1,
		},
	}

	runTests(t, tests)
}
