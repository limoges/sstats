package sstats

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	newSum := func(v float64) *Sum {
		s := Sum(v)
		return &s
	}
	tests := []statTest{
		{
			Init:     newSum(0),
			Values:   nil,
			Valid:    true,
			Expected: 0.0,
		},
		{
			Init:     newSum(0),
			Values:   []float64{1},
			Valid:    true,
			Expected: 1,
		},
		{
			Init:     newSum(0),
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 15,
		},
		{
			Init:     newSum(0),
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 3,
		},
		{
			Init:     newSum(0),
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 2,
		},
		{
			Init:     newSum(0),
			Values:   []float64{math.SmallestNonzeroFloat64, 10},
			Valid:    true,
			Expected: 10,
		},
		{
			Init:     newSum(0),
			Values:   []float64{math.Inf(1), -math.MaxFloat64},
			Valid:    true,
			Expected: math.Inf(1),
		},
		{
			Init:     newSum(0),
			Values:   []float64{math.Inf(-1), math.MaxFloat64},
			Valid:    true,
			Expected: math.Inf(-1),
		},
	}

	runTests(t, tests)
}
