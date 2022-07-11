package leetcode148

import "github.com/lucky51/algo/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *common.ListNode) *common.ListNode {
	return sort(head, nil)
}

func sort(head, tail *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		if fast.Next != tail {
			fast = fast.Next.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}
func merge(head1, head2 *common.ListNode) *common.ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	if head1.Val < head2.Val {
		head1.Next = merge(head1.Next, head2)
		return head1
	} else {
		head2.Next = merge(head2.Next, head1)
		return head2
	}
}
