package sstats

import "fmt"

type Count int

func (c *Count) Push(values ...float64) {
	new_count := int(*c) + len(values)
	*c = Count(new_count)
}

func (c Count) Value() (float64, bool) {
	return float64(c), true
}

func (c Count) String() string {
	return fmt.Sprintf("Count:%v", int(c))
}
