package leetcode16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, best := len(nums), math.MaxInt32
	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b, c := a+1, n-1
		for b < c {
			if nums[a]+nums[b]+nums[c] == target {
				return target
			}
			var sub, curr = nums[a] + nums[b] + nums[c] - target, best - target
			if sub < 0 {
				sub = -sub
			}
			if curr < 0 {
				curr = -curr
			}
			if sub < curr {
				best = nums[a] + nums[b] + nums[c]
			}

			if nums[a]+nums[b]+nums[c] < target {
				b0 := b + 1
				// 不能这么写，这么写就是 b+1 和 b+2比较了，而不是 b和 b+1 比较
				// for ; b0+1 < c && nums[b0+1] == nums[b0]; b0++ {
				// }
				for ; b0+1 < c && nums[b0] == nums[b]; b0++ {
				}
				b = b0
			} else {
				c0 := c - 1
				for ; c0-1 > b && nums[c0] == nums[c]; c0-- {
				}
				c = c0
			}
		}
	}
	return best
}
