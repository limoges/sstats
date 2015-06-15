package sstats

type maxHeap []float64

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Peek() interface{}   { return h[0] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(float64)) }
func (h *maxHeap) Pop() interface{} {
	p := *h
	n := len(p)
	v := p[n-1]
	*h = p[0 : n-1]
	return v
}
