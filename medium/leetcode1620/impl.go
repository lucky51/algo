package leetcode1620

import "math"

// bestCoordinate 网络信号最好的坐标
// 这道题计算范围是所有塔覆盖的所有坐标，而非各塔单个坐标
func bestCoordinate(towers [][]int, radius int) []int {
	maxX, maxY := 0, 0
	var cx, cy, maxQuality int
	for _, t := range towers {
		maxX = max(maxX, t[0])
		maxY = max(maxY, t[1])
	}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			quality := 0
			for _, t := range towers {
				d := (x-t[0])*(x-t[0]) + (y-t[1])*(y-t[1])
				if d <= radius*radius {
					quality += int(float64(t[2]) / (1 + math.Sqrt(float64(d))))
				}
			}
			// 大于才替换，就能保证相同信号的坐标保持最小字典序
			// 坐标 左 ——> 右 ，下——>上 在加上下边的这个条件可以保证字典序排列的最小坐标
			if quality > maxQuality {
				cx, cy, maxQuality = x, y, quality
			}
		}
	}
	return []int{cx, cy}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
