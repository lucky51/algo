package leetcode237

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteNode 删除链表中的节点
func deleteNode(node *ListNode) {
	// 这道题不能按常规的思考方式进行解题，因为你无法得到上一个节点，只能用值替换
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
