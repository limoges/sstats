package sstats

import "fmt"

type Minimum struct {
	set bool
	min float64
}

func (m *Minimum) Push(values ...float64) {
	for _, v := range values {
		if !m.set || v < m.min {
			m.set = true
			m.min = v
		}
	}
}

func (m *Minimum) Value() (float64, bool) {
	return m.min, m.set
}

func (m Minimum) String() string {
	if m.set {
		return fmt.Sprintf("Min:%v", m.min)
	}
	return "Min:-"
}
