package sstats

import "testing"

func TestStddev(t *testing.T) {
	tests := []statTest{
		{
			Init:     &Stddev{},
			Values:   nil,
			Valid:    false,
			Expected: 0.0,
		},
		{
			Init:     &Stddev{},
			Values:   []float64{1},
			Valid:    true,
			Expected: 0,
		},
		{
			Init:     &Stddev{},
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 1.870828693,
		},
		{
			Init:     &Stddev{},
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     &Stddev{},
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 1.414213562,
		},
	}

	runTests(t, tests)
}
