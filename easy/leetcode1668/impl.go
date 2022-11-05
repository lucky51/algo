package leetcode1668

// maxRepeating  最大重复子字符串
func maxRepeating(sequence string, word string) int {
	ans := 0
	n, m := len(sequence), len(word)
	f := make([]int, n)
	for i := m - 1; i < n; i++ {
		if sequence[i-m+1:i+1] == word {
			if i == m-1 {
				f[i] = 1
			} else {
				f[i] = f[i-m] + 1
			}
			if f[i] > ans {
				ans = f[i]
			}
		}
	}
	return ans
}
