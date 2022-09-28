package leetcodekthmagicnumber

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

var factories = []int{3, 5, 7}

// getKthMagicNumber 面试题17.09.第k个数
func getKthMagicNumber(k int) int {
	seen := map[int]struct{}{1: {}}
	h := &hp{sort.IntSlice{1}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == k {
			return x
		}
		for _, f := range factories {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}
