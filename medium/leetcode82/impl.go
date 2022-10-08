package leetcode82

import . "github.com/lucky51/algo/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head, Val: -101}
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		x := cur.Next.Val
		if cur.Next.Val == cur.Next.Next.Val {
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
