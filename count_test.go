package sstats

import "testing"

func TestCount(t *testing.T) {
	tests := []statTest{
		{
			Init:     NewCount(0),
			Values:   nil,
			Valid:    true,
			Expected: 0,
		},
		{
			Init:     NewCount(0),
			Values:   []float64{0, 2, 1, 3, 4, 5},
			Valid:    true,
			Expected: 6,
		},
		{
			Init:     NewCount(2),
			Values:   []float64{0, 2, 1},
			Valid:    true,
			Expected: 5,
		},
		{
			Init:     NewCount(-1),
			Values:   []float64{0, 2},
			Valid:    true,
			Expected: 1,
		},
	}

	runTests(t, tests)

	got := NewCount(0).String()
	expected := "Count:0"
	if got != expected {
		t.Errorf("got %s, expected %s", got, expected)
	}
}
