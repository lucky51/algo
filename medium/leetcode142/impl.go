package leetcode142

import "github.com/lucky51/algo/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *common.ListNode) *common.ListNode {
	nodeMap := map[*common.ListNode]struct{}{}
	ans := head
	for ans != nil {
		if _, ok := nodeMap[ans]; ok {
			return ans
		}
		nodeMap[ans] = struct{}{}
		ans = ans.Next
	}
	return ans
}
