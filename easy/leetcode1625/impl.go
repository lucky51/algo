package leetcode1625

// decrypt 拆炸弹  ,自己的解法 ，官方题解为 滑动窗口
func decrypt(code []int, k int) []int {
	ans := []int{}
	n := len(code)
	if k == 0 {
		return make([]int, n)
	}
	for i := range code {
		sum := 0
		j := i
		cnt := k
		if k < 0 {
			cnt = -k
		}
		for cnt > 0 {
			if k > 0 {
				j = (j + 1) % n
			} else {
				j = (j - 1 + n) % n
			}
			sum += code[j]
			cnt--
		}
		ans = append(ans, sum)
	}
	return ans
}

// decrypt1 滑动窗口 官方题解
func decrypt1(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}
	code = append(code, code...)
	l, r := 1, k
	if k < 0 {
		// 这里的k是负数了所以+
		l = n + k
		r = n - 1
	}
	sum := 0
	for _, v := range code[l : r+1] {
		sum += v
	}
	for i := range ans {
		ans[i] = sum
		sum -= code[l]
		sum += code[r+1]
		l++
		r++
	}
	return ans
}
