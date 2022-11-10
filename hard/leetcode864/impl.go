package leetcode864

import "unicode"

// TODO这道题没有理解  ，题解中使用了二进制来进行状态压缩，二进制位运算没有理解
/*
算题群里大佬解析了下
简化思想，假设只有两种钥匙，则可以用4个二进制数表示钥匙拥有的状态。
00 ---- 没有钥匙
01 ---- 有1号钥匙，没有2号钥匙
10 ---- 有2号钥匙，没有1号钥匙
11 ---- 拥有两把钥匙

如何从“没有钥匙”转化为“拥有第二把钥匙”呢？
10 = 00 | 10 = 00 | (1<<1)

这里或运算，就代表这一位的钥匙从无到有了。（而且1|1=1，所以不怕重复运算）
1<<1，就是1左移一次，正好是10，代表第二把钥匙。
*/
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// shortestPathAllKeys  获取所有钥匙的最短路径
func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	sx, sy := 0, 0
	keyToIdx := map[rune]int{}
	for i, row := range grid {
		for j, c := range row {
			if c == '@' {
				sx, sy = i, j
			} else if unicode.IsLower(c) {
				if _, ok := keyToIdx[c]; !ok {
					keyToIdx[c] = len(keyToIdx)
				}
			}
		}
	}

	dist := make([][][]int, m)
	for i := range dist {
		dist[i] = make([][]int, n)
		for j := range dist[i] {
			dist[i][j] = make([]int, 1<<len(keyToIdx))
			for k := range dist[i][j] {
				dist[i][j][k] = -1
			}
		}
	}
	dist[sx][sy][0] = 0
	q := [][3]int{{sx, sy, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, mask := p[0], p[1], p[2]
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] != '#' {
				c := rune(grid[nx][ny])
				if c == '.' || c == '@' {
					if dist[nx][ny][mask] == -1 {
						dist[nx][ny][mask] = dist[x][y][mask] + 1
						q = append(q, [3]int{nx, ny, mask})
					}
				} else if unicode.IsLower(c) {
					t := mask | 1<<keyToIdx[c]
					if dist[nx][ny][t] == -1 {
						dist[nx][ny][t] = dist[x][y][mask] + 1
						if t == 1<<len(keyToIdx)-1 {
							return dist[nx][ny][t]
						}
						q = append(q, [3]int{nx, ny, t})
					}
				} else {
					idx := keyToIdx[unicode.ToLower(c)]
					if mask>>idx&1 > 0 && dist[nx][ny][mask] == -1 {
						dist[nx][ny][mask] = dist[x][y][mask] + 1
						q = append(q, [3]int{nx, ny, mask})
					}
				}
			}
		}
	}
	return -1
}

//作者：力扣官方题解
//链接：https://leetcode.cn/problems/shortest-path-to-get-all-keys/solutions/1959739/huo-qu-suo-you-yao-chi-de-zui-duan-lu-ji-uu5m/
//来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
