package leetcode816

import "fmt"

// ambiguousCoordinates 模糊坐标
func ambiguousCoordinates(s string) []string {
	f := func(i, j int) []string {
		res := []string{}
		// 枚举开始到结束中，“.” 可以插入的位置
		for k := 1; k <= j-i; k++ {
			l, r := s[i:i+k], s[i+k:j]
			// 插入“.”的要求：左侧要么为 0要么开头不能为0，右侧的最后一个数字不能为0
			ok := (l == "0" || l[0] != '0') && (r == "" || r[len(r)-1] != '0')
			if ok {
				t := ""
				if k < j-i {
					t = "."
				}
				res = append(res, l+t+r)
			}
		}
		return res
	}
	n := len(s)
	ans := []string{}
	for i := 2; i < n-1; i++ {
		for _, x := range f(1, i) {
			for _, y := range f(i, n-1) {
				ans = append(ans, fmt.Sprintf("(%s, %s)", x, y))
			}
		}
	}
	return ans
}
