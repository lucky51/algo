package leetcode347

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IHeap1{}
	heap.Init(h)
	heap.Push(h, 5)
	heap.Push(h, 6)
	heap.Push(h, 7)
	heap.Push(h, 1)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}

type IHeap1 []int

func (h IHeap1) Len() int           { return len(h) }
func (h IHeap1) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap1) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IHeap1) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
