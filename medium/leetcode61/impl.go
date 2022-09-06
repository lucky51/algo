package leetcode61

import . "github.com/lucky51/algo/common"

// rotateRight 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	itor := head
	for itor.Next != nil {
		n++
		itor = itor.Next
	}
	add := n - (k % n)
	if add == n {
		return head
	}
	// 把末尾元素指向头指针，形成环
	itor.Next = head
	for add > 0 {
		itor = itor.Next
		add--
	}
	head = itor.Next
	itor.Next = nil
	return head
}
