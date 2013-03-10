package sstats

import (
	"fmt"
	"strings"
)

type Statistic interface {
	Push(values ...float64)
	Value() (float64, bool)
}

type SimpleStats struct {
	stats []Statistic
}

func NewSimpleStats(stats ...Statistic) SimpleStats {
	return SimpleStats{
		stats: stats,
	}
}

func (s SimpleStats) Push(values ...float64) {
	for _, v := range s.stats {
		v.Push(values...)
	}
}

func (s SimpleStats) Values() []float64 {
	var values []float64
	for _, v := range s.stats {
		if val, ok := v.Value(); ok {
			values = append(values, val)
		}
	}
	return values
}

func (s SimpleStats) String() string {
	str := make([]string, len(s.stats))
	for i, v := range s.stats {
		str[i] = fmt.Sprint(v)
	}
	return fmt.Sprint(strings.Join(str, ","))
}
