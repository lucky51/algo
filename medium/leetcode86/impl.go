package leetcode86

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// partition 分隔链表
func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
