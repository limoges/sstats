package sstats

import "fmt"

type Maximum struct {
	set bool
	max float64
}

func (m *Maximum) Push(values ...float64) {
	for _, v := range values {
		if !m.set || v > m.max {
			m.set = true
			m.max = v
		}
	}
}

func (m Maximum) Value() (float64, bool) {
	return m.max, m.set
}

func (m Maximum) String() string {
	if m.set {
		return fmt.Sprintf("Max:%v", m.max)
	}
	return "Max:-"
}
