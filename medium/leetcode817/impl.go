package leetcode817

import . "github.com/lucky51/algo/common"

func numComponents(head *ListNode, nums []int) int {
	// 题目不容易理解，在链表中计算有多少组件的起始位置
	memo := map[int]struct{}{}
	for _, n := range nums {
		memo[n] = struct{}{}
	}
	ans := 0
	for inSide := false; head != nil; head = head.Next {
		if _, ok := memo[head.Val]; !ok {
			inSide = false
		} else if !inSide {
			inSide = true
			ans++
		}
	}
	return ans
}
