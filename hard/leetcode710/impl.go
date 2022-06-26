package leetcode710

import "math/rand"

// Solution
// b2w 用边界左侧的黑名单元素映射到右侧的白名单元素
type Solution struct {
	b2w   map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	// bound 这里边指边界，因为是数的范围
	bound := n - m
	black := map[int]bool{}
	// 寻找右侧的黑名单元素
	for _, b := range blacklist {
		// 0 1 2 3 4 5 6     传入 n=7，m=3   ,那么边界 bound = 7-3  就是 b =bound
		if b >= bound {
			black[b] = true
		}
	}
	// 映射的长度应该等于，左侧的和名单元素数量
	b2w := make(map[int]int, m-len(black))
	// 记录一个临时的边界值，用来推导右侧中白名单元素的位置
	w := bound
	// 查找左侧黑名单元素
	for _, b := range blacklist {
		if b < bound {
			// 找到右侧第一个白名单元素
			for black[w] {
				w++
			}
			// 将当前找到的左侧黑名单元素映射到bound边界找到的第一个白名单元素
			b2w[b] = w
			w++
		}
	}
	return Solution{b2w, bound}
}

// Pick 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
func (this *Solution) Pick() int {
	x := rand.Intn(this.bound)
	if this.b2w[x] > 0 {
		return this.b2w[x]
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
