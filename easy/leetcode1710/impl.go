package leetcode1710

import "sort"

/*
	这里的truckSize是最大装载箱子数量
	贪心算法，每次取最大的单元格数
*/

// maximumUnits 卡车上的最大单元数
func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	// 先按照单元格数量逆序排列
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	for _, t := range boxTypes {
		// 如果当前箱子数量大于最大数量，计算剩余的单元格数量
		if t[0] >= truckSize {
			ans += truckSize * t[1]
			break
		}
		truckSize -= t[0]
		ans += t[0] * t[1]
	}
	return
}
