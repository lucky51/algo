package leetcode757

import (
	"sort"
)

// intersectionSizeTwo 设置交集大小至少等于2
func intersectionSizeTwo(intervals [][]int) (ans int) {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	n, m := len(intervals), 2
	// 这里的vals只是用来记录大于左边界的区间数，里边的值没有实际含义  可以替换成 map[int]int 记录一下重叠的区间数，改成map就会好理解一些，这里强调的是数量而不是内容值
	// 这里边比较难理解的是，如果区间的右边界大于目标区间的左边界s[i,n]，那么 si 一定在目标区间内，如果其中的一个区间在另外两个区间内有交集，则可以跳过，这里边题目的 intervals[i] 最小长度也要等于2
	// 分别代表左右边界，这块不怎么好理解
	vals := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		for j, k := intervals[i][0], len(vals[i]); k < m; k++ {
			ans++
			for p := i - 1; p >= 0 && intervals[p][1] >= j; p-- {
				vals[p] = append(vals[p], j)
			}
			j++
		}
	}
	return
}
