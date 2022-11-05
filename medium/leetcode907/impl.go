package leetcode907

// sumSubarrayMins 子数组的最小值之和
// 官方题解没有看懂，参考题解
// https://leetcode.cn/problems/sum-of-subarray-minimums/solutions/1930857/gong-xian-fa-dan-diao-zhan-san-chong-shi-gxa5/
// 大概意思懂了，但是具体的边界还是没有太清楚
// 这个题解最通俗的理解方式，后续优化过程没有仔细看
func sumSubarrayMins(arr []int) (ans int) {
	n := len(arr)
	// left 为左边严格小于 arr[i]的最近元素位置 (不存在时候为-1，)
	left := make([]int, n)
	st := []int{-1}
	for i, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	// 右边界 right为右侧小于等于arr[i]的最近元素位置(不存在是为n)
	right := make([]int, n)
	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		x := arr[i]
		// 这块算的都是边界 开区间 () ,不是算最小值，而是最小值的下一个较大值
		for len(st) > 1 && arr[st[len(st)-1]] > x {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	for i, x := range arr {
		ans += x * (i - left[i]) * (right[i] - i) // 累加贡献  ，数量 * 最小值
	}
	return ans % (1e9 + 7)
}
