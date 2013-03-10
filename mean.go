package sstats

type Mean struct {
	n   int
	sum float64
}

func (m *Mean) Push(values ...float64) {
	m.n += len(values)
	for _, v := range values {
		m.sum += v
	}
}

func (m Mean) Value() (float64, bool) {
	if m.n == 0 {
		return 0.0, false
	}
	return m.sum / float64(m.n), true
}
