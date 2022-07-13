package leetcode735

// asteroidCollision 行星碰撞
func asteroidCollision(asteroids []int) []int {
	// 这里有两个身份 一个 是  aster ,另一个是 alive
	// 解法栈模拟
	// 定义栈
	st := []int{}
	for _, aster := range asteroids {
		// 定义当前行星是否还活着（未爆炸）
		alive := true
		for alive && aster < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			// 仅有当前行星比栈顶反向的行星大才能活着
			alive = st[len(st)-1] < -aster
			// 当前行星大于等于栈顶行星时，栈顶行星爆炸，移除栈
			if st[len(st)-1] <= -aster {
				st = st[:len(st)-1]
			}
		}
		// 如果当前行星依然活着，则入栈
		if alive {
			st = append(st, aster)
		}
	}
	return st
}
