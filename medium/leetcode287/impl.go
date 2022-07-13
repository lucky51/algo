package leetcode287

import "fmt"

// 寻找重复数 ，官方解法三，快慢指针
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
		fmt.Println(slow, fast)
	}
	fmt.Println("end loop", slow, fast)
	slow = 0
	/*
		如果慢指针从起点出发，快指针从相遇位置出发，每次两个指针都移动一步，
		则慢指针走了 a 步之后到达环的入口，快指针在环里走了 k−1 圈之后又走了 c 步，
		由于从相遇位置继续走 c 步即可回到环的入口，因此快指针也到达环的入口。两个指针在环的入口相遇，相遇点就是答案。
	*/
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 寻找重复数 官方解法一 ，二分查找
// 易读的题解 https://leetcode.cn/problems/find-the-duplicate-number/solution/mei-kan-dong-er-fen-cha-zhao-ti-jie-de-x-0ixw/
func findDuplicateCnt(nums []int) int {
	ans := -1
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		cnt := 0
		mid := (lo + hi) >> 1
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			lo = mid + 1
		} else {
			hi = mid - 1
			ans = mid
		}
	}
	return ans
}
