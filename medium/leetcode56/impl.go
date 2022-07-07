package leetcode56

import "sort"

// merge 合并区间
func merge(intervals [][]int) [][]int {
	//1. 先按左边界排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	// 定义结果集
	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		// 如果当前区间的左边界大于最后一个的右边界，则说明没有任何交集
		if intervals[i][0] > last[1] {
			ans = append(ans, intervals[i])
		} else {
			//反之，更新最后一个区间的右边界
			last[1] = max(intervals[i][1], last[1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
