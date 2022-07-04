package leetcode556

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"unsafe"
)

// nextGreaterElement 下一个更大元素 III ,参考题解
// https://leetcode.cn/problems/next-greater-element-iii/solution/by-lfool-vi69/
func nextGreaterElement(n int) (ans int) {
	chars := []byte(strconv.Itoa(n))
	fmt.Println(chars)
	// 向前找一个数字  chars[i] > chars[i-1]的 升序
	for i := len(chars) - 1; i > 0; i-- {
		if chars[i] > chars[i-1] {
			temp := chars[i:]
			sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
			fmt.Println("temp", temp)
			fmt.Println(chars)
			for j := i; j < len(chars); j++ {
				if chars[j] > chars[i-1] {
					chars[i-1], chars[j] = chars[j], chars[i-1]
					fmt.Println(chars)
					newStr := *(*string)(unsafe.Pointer(&chars))
					ans, _ := strconv.ParseInt(newStr, 10, 64)
					if ans > math.MaxInt32 {
						return -1
					} else {
						return int(ans)
					}
				}
			}
		}
	}
	return -1
}
