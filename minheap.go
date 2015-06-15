package sstats

type minHeap []float64

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h minHeap) Peek() interface{}   { return h[0] }
func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(float64)) }
func (h *minHeap) Pop() interface{} {
	p := *h
	n := len(p)
	v := p[n-1]
	*h = p[0 : n-1]
	return v
}
