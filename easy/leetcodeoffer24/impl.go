package leetcodeoffer24

import . "github.com/lucky51/algo/common"

// reverseList 翻转链表
// 这个翻转顺序总是记不住
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
