package leetcode827

// largestIsland 最大人工岛 TODO:这道题还要再写一遍，抄了一半
func largestIsland(grid [][]int) (ans int) {
	// 上下左右四个方向
	dir4 := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, t := len(grid), 0
	// 用来记录某个点是否被访问过
	tag := make([][]int, n)
	// 记录岛屿和对应的面积
	area := map[int]int{}
	for i := range tag {
		tag[i] = make([]int, n)
	}
	// 深度优先遍历计算岛屿面积
	var dfs func(int, int)
	dfs = func(i, j int) {
		tag[i][j] = t
		area[t]++
		for _, d := range dir4 {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] > 0 && tag[x][y] == 0 {
				dfs(x, y)
			}
		}
	}
	// 对每个大于的点进行深度优先遍历，以来计算每个岛屿的面积
	for i, row := range grid {
		for j, x := range row {
			if x > 0 && tag[i][j] == 0 {
				// 用来标识同一岛屿
				t = i*n + j + 1
				dfs(i, j)
				ans = max(ans, area[t])
			}
		}
	}
	// 枚举可以添加陆地的位置
	for i, row := range grid {
		for j, x := range row {
			if x == 0 {
				// 默认为1 ，因为当前陆地的大小
				newArea := 1
				conn := map[int]bool{0: true}
				for _, d := range dir4 {
					x, y := i+d[0], j+d[1]
					if x >= 0 && x < n && y >= 0 && y < n && !conn[tag[x][y]] {
						newArea += area[tag[x][y]]
						conn[tag[x][y]] = true
					}
				}
				ans = max(ans, newArea)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
