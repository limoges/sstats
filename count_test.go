package sstats

import "testing"

func TestCount(t *testing.T) {
	count := Count(0)
	count1 := Count(0)
	count2 := Count(2)
	count3 := Count(-1)
	tests := []statTest{
		{
			Init:     &count,
			Values:   nil,
			Valid:    true,
			Expected: 0,
		},
		{
			Init:     &count1,
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 6,
		},
		{
			Init:     &count2,
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 5,
		},
		{
			Init:     &count3,
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 1,
		},
	}

	runTests(t, tests)
}
