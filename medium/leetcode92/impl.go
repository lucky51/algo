package leetcode92

import . "github.com/lucky51/algo/common"

// reverseBetween 反转链表II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 构造虚拟节点
	dummyNode := &ListNode{Val: -1, Next: head}
	prev := dummyNode
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	rightNode := prev
	// 遍历 right-left+1到达右节点
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	// 切断出一个新的链表用于反转
	leftNode := prev.Next
	curr := rightNode.Next
	// 切断链接
	prev.Next = nil
	rightNode.Next = nil
	// 翻转链表
	reverseLinkedList(leftNode)
	// 链接翻转后的链表
	prev.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

// reverseLinkedList 单链表反转
func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}
