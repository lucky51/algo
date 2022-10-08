package leetcode239

import (
	"container/heap"
	"sort"
)

/*
 使用优先队列(大顶堆)来对窗口内的元素进行排序，大的放在上边
 每次遍滑动窗口就能得出当前的最大值
*/

var a []int

// 定义优先队列，大顶堆
type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// maxSlidingWindow 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	// 优先队列根据传入数组进行排序
	a = nums
	n := len(nums)
	h := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[h.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(h, i)
		// i-k窗口范围之外的元素要删掉
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[h.IntSlice[0]])
	}
	return ans
}
