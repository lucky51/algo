package leetcode347

import "container/heap"

// topKFrequent 前K个高频元素
// 利用小根堆的特性，对每个数字出现频次数进行从小到大的排序，从而实现第几的效果
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, v := range nums {
		occurrences[v]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, val := range occurrences {
		// 这个Push 和 Pop一定要借助 heap去调用，否则没效果
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() any {
	old := *h
	x := old[len(*h)-1]
	*h = old[:len(*h)-1]
	return x
}
