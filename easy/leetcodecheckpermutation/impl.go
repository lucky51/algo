package leetcodecheckpermutation

// CheckPermutation 面试题01.02 判定是否互为字符重排
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	n := len(s1)
	memo := make(map[byte]int, n)
	for i := 0; i < len(s1); i++ {
		memo[s1[i]]++
		memo[s2[i]]--
	}
	for _, v := range memo {
		if v != 0 {
			return false
		}
	}
	return true
}
