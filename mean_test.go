package sstats

import "testing"

func TestMean(t *testing.T) {
	tests := []statTest{
		{
			Init:     &Mean{},
			Values:   nil,
			Valid:    false,
			Expected: 0.0,
		},
		{
			Init:     &Mean{},
			Values:   []float64{1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Mean{},
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 2.5,
		},
		{
			Init:     &Mean{},
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Mean{},
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 1,
		},
	}

	runTests(t, tests)
}
