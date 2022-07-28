package leetcode1331

import "sort"

// arrayRankTransform 数组序号转换
func arrayRankTransform(arr []int) []int {
	a := append([]int(nil), arr...)
	sort.Ints(a)
	rank := map[int]int{}
	for _, v := range a {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1
		}
	}
	for i, v := range rank {
		arr[i] = rank[v]
	}
	return arr
}
