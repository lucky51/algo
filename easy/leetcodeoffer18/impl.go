package leetcodeoffer18

import . "github.com/lucky51/algo/common"

// deleteNode 剑指 Offer 18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head}
	var prev, curr *ListNode
	curr = head
	prev = dummyNode
	for curr != nil {
		if curr.Val == val {
			break
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = prev.Next.Next
	return dummyNode.Next
}
