package leetcode749

type pair struct{ x, y int }

// 四个方向
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// containVirus 隔离病毒
func containVirus(isInfected [][]int) (ans int) {
	m, n := len(isInfected), len(isInfected[0])
	for {
		neighbors := []map[pair]struct{}{}
		firewalls := []int{}
		for i, row := range isInfected {
			// 寻找感染坐标位置
			for j, infected := range row {
				if infected != 1 {
					continue
				}
				// 广度优先遍历
				q := []pair{{i, j}}
				// 记录当前感染区域中的所有相邻非感染节点
				neighbor := map[pair]struct{}{}
				// 记录当前区域的防火墙 ，这里的idx要从1开始，这样 -(1)才会区别于原始的值 0,1
				firewall, idx := 0, len(neighbors)+1
				// 修改当前节点的值，和处理过的区分开来
				row[j] = -idx
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					// 寻找当前被坐标四周所有别感染的坐标，然后入队
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n {
							if isInfected[x][y] == 1 {
								q = append(q, pair{x, y})
								isInfected[x][y] = -idx
							} else if isInfected[x][y] == 0 {
								// 如果是非感染坐标，则可以成为防火墙
								firewall++
								neighbor[pair{x, y}] = struct{}{}
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}
		if len(neighbors) == 0 {
			break
		}
		idx := 0
		// 找到本次感染
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}
		// 更新结果
		ans += firewalls[idx]
		for _, row := range isInfected {
			for j, v := range row {
				if v < 0 {
					// 这里的idx是从0开始的，而在矩阵中记录的值是从(-)1开始的,所以要 减1
					if v != -idx-1 {
						// 其余的更新为 1，已方便下一次遍历已经扩散的新矩阵
						row[j] = 1
					} else {
						//把当前最大区域的值更新为 2,下一次遍历排除已经封闭的区域
						row[j] = 2
					}
				}
			}
		}
		// 这里要把所有非最大感染的所有区域相邻的坐标都改为1 ，因为下一次循环这些区域会被感染
		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}
		// 如果只有一个，则代表只存在一个我们添加了防火墙的区域，则无需在处理
		if len(neighbors) == 1 {
			break
		}
	}
	return
}
