package leecode2250

import (
	"sort"
)

// countRectangles 统计包含每个点的矩形数目
func countRectangles(rectangles [][]int, points [][]int) []int {
	// 让矩形按照y进行从大大小排列
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i][1] > rectangles[j][1] })
	// 记录一下每一个点的下标位置
	for pIndex := range points {
		points[pIndex] = append(points[pIndex], pIndex)
	}
	// 让点按照y进行从大到小排列
	sort.Slice(points, func(i, j int) bool { return points[i][1] > points[j][1] })
	// xs 的长度不是 make([]int, len(rectangles))
	// i, xs, ans := 0, make([]int, len(rectangles)), make([]int, len(points))
	i, xs, ans := 0, []int{}, make([]int, len(points))
	for p := 0; p < len(points); p++ {
		start := i
		for ; i < len(rectangles) && rectangles[i][1] >= points[p][1]; i++ {
			xs = append(xs, rectangles[i][0])
		}
		// 如果有新元素添加，则排序
		if start != i {
			sort.Ints(xs)
		}
		//i代表当前符合y条件的所有矩形数，查找当前xs中符合点x的下标，获取大于等于该点x坐标的所有数量就是这个点符合所有矩形的数量
		ans[points[p][2]] = i - sort.SearchInts(xs, points[p][0])
	}
	return ans
}
