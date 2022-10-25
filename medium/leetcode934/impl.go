package leetcode934

/*
 通过两次广度优先遍历
 第一次标记岛屿
 第二次将已经标记的岛屿向外扩散
*/

// shortestBridge 最短的桥
func shortestBridge(grid [][]int) (step int) {
	n := len(grid)
	dir := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	q := [][2]int{}
	for i, row := range grid {
		for j, c := range row {
			if c == 0 {
				continue
			}
			island := [][2]int{}
			q = append(q, [2]int{i, j})
			grid[i][j] = -1
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				island = append(island, p)
				for _, d := range dir {
					x, y := p[0]+d[0], p[1]+d[1]
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 1 {
						grid[x][y] = -1
						q = append(q, [2]int{x, y})
					}
				}
			}

			q = island
			for {
				tmp := q
				q = nil
				for _, p := range tmp {
					for _, d := range dir {
						x, y := p[0]+d[0], p[1]+d[1]
						if x >= 0 && x < n && y >= 0 && y < n {
							if grid[x][y] == -1 {
								continue
							}
							if grid[x][y] == 1 {
								return
							} else if grid[x][y] == 0 {
								grid[x][y] = -1
								q = append(q, [2]int{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return
}
