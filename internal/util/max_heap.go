package util

type HeapItem struct {
	Priority int
	I        int
	J        int
}

type IntMaxHeap []HeapItem

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i].Priority > h[j].Priority }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(HeapItem))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
