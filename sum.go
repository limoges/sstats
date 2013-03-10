package sstats

import "fmt"

type Sum float64

func (s *Sum) Push(values ...float64) {
	// To minimize the error, one should sort the values in increasing order
	// before summing the floats
	for _, v := range values {
		new_sum := float64(*s) + v
		*s = Sum(new_sum)
	}
}

func (s Sum) Value() (float64, bool) {
	return float64(s), true
}

func (s Sum) String() string {
	return fmt.Sprintf("Sum:%v", float64(s))
}
