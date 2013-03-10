package sstats

import (
	"reflect"
	"testing"
)

type mockStat struct {
	pushCalls  int
	values     []float64
	valueCalls int
}

func (s *mockStat) Push(values ...float64) {
	s.pushCalls++

	if len(values) < 1 {
		return
	}

	s.values = append(s.values, values...)
}

func (s *mockStat) Value() (float64, bool) {
	s.valueCalls++
	return 0.0, true
}

func TestSimpleStats(t *testing.T) {
	tests := []struct {
		calls      [][]float64
		expected   []float64
		valueCalls int
	}{
		{
			calls: [][]float64{
				[]float64{1, 2, 3},
				[]float64{4, 7, 8},
				nil,
				[]float64{2, 3, 9}},
			expected:   []float64{1, 2, 3, 4, 7, 8, 2, 3, 9},
			valueCalls: 1,
		},
	}

	for _, test := range tests {
		mock := mockStat{}
		sstats := NewSimpleStats(&mock)
		for _, vals := range test.calls {
			sstats.Push(vals...)
		}
		for i := 0; i < test.valueCalls; i++ {
			sstats.Values()
		}
		if mock.pushCalls != len(test.calls) {
			t.Errorf("got %v, expected %v", mock.pushCalls, len(test.calls))
		}
		if mock.valueCalls != test.valueCalls {
			t.Errorf("got %v, expected %v", mock.valueCalls, test.valueCalls)
		}
		if !reflect.DeepEqual(mock.values, test.expected) {
			t.Errorf("got %v, expected %v", mock.values, test.expected)
		}
	}
}
