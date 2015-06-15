package sstats

import "fmt"

type Sum float64

func NewSum(v float64) *Sum {
	s := Sum(v)
	return &s
}

func (s *Sum) Push(values ...float64) {
	// To minimize the error, one should sort the values in increasing order
	// before summing the floats
	for _, v := range values {
		*s += Sum(v)
	}
}

func (s Sum) Value() (float64, bool) {
	return float64(s), true
}

func (s Sum) String() string {
	return fmt.Sprintf("Sum:%f", s)
}
