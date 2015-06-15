package sstats

import "fmt"

type Count int

func NewCount(v int) *Count {
	c := Count(v)
	return &c
}

func (c *Count) Push(values ...float64) {
	*c += Count(len(values))
}

func (c Count) Value() (float64, bool) {
	return float64(c), true
}

func (c Count) String() string {
	return fmt.Sprintf("Count:%d", c)
}
