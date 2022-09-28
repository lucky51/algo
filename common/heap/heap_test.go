package heap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

func TestSmallHeap(t *testing.T) {
	hp := hp{sort.IntSlice{1}}
	heap.Push(&hp, 5)
	heap.Push(&hp, 4)
	heap.Push(&hp, 1)
	heap.Push(&hp, 99)
	fmt.Println(heap.Pop(&hp))

}
