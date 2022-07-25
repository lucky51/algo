package leetcode1184

// distanceBetweenBusStops 公交站间的距离 TODO： leetcode 打卡题目
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	// 这道题可以使用一次遍历，因为遍历所有的distance就是从start到distance，就包括了正向和反向遍历
	// 0, 1, 2, 3,  4,   5,    6,    7,    8,
	//         start           end
	//我们只需关注start 到end之前的距离，然后剩下的就是反向，求出两个和值，比较一下最小的就是答案
	if start > destination {
		start, destination = destination, start
	}
	sum1, sum2 := 0, 0
	for i, d := range distance {
		// 这里的边界需要注意下，distance[i]表示的是i到 (i+1) % n 之间的距离
		if i >= start && i < destination {
			sum1 += d
		} else {
			sum2 += d
		}
	}
	if sum1 < sum2 {
		return sum1
	}
	return sum2
}
