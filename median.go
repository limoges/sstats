package sstats

import (
	"container/heap"
	"fmt"
)

type minmaxheap interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Median struct {
	lower maxHeap
	upper minHeap
}

func NewMedian() *Median {
	m := Median{}
	heap.Init(&m.lower)
	heap.Init(&m.upper)
	return &m
}

func (m Median) len() int {
	return m.lower.Len() + m.upper.Len()
}

func (m *Median) Push(values ...float64) {
	for _, v := range values {
		m.add(v)
		m.rebalance()
	}
}

func (m Median) Value() (float64, bool) {
	if m.len() < 1 {
		return 0.0, false
	}

	bigger, smaller := m.heapsByLen()
	if bigger.Len() == smaller.Len() {
		return (bigger.Peek().(float64) + smaller.Peek().(float64)) / 2.0, true
	} else {
		return bigger.Peek().(float64), true
	}
}

func (m *Median) add(x float64) {
	if m.lower.Len() == 0 || x < m.lower.Peek().(float64) {
		heap.Push(&m.lower, x)
	} else {
		heap.Push(&m.upper, x)
	}
}

func (m *Median) rebalance() {
	bigger, smaller := m.heapsByLen()
	if bigger.Len()-smaller.Len() >= 2 {
		heap.Push(smaller, heap.Pop(bigger))
	}
}

func (m Median) heapsByLen() (bigger, smaller minmaxheap) {
	if m.lower.Len() > m.upper.Len() {
		return &m.lower, &m.upper
	}
	return &m.upper, &m.lower
}

func (m Median) String() string {
	median, valid := m.Value()
	if valid {
		return fmt.Sprintf("Median:%v", median)
	}
	return fmt.Sprintf("Median:-")
}
