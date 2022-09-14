package leetcode1619

import "sort"

// trimMean 删除某些元素后的数组均值
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n:=len(arr)
	sum:=0
	for _,num:=range arr[n/20:n*19/20] {
		sum +=num
	}
	return float64(sum * 10) / float64(n * 9)
}
