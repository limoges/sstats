package sstats

import (
	"fmt"
	"math"
)

type Stddev struct {
	variance Variance
}

func (s *Stddev) Push(values ...float64) {
	s.variance.Push(values...)
}

func (s Stddev) Value() (float64, bool) {
	variance, valid := s.variance.Value()
	return math.Sqrt(variance), valid
}

func (s Stddev) String() string {
	stddev, valid := s.Value()
	if valid {
		return fmt.Sprintf("Stddev:%v", stddev)
	}
	return fmt.Sprintf("Stddev:-")
}
