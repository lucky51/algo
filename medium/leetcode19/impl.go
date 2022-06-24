package leetcode19

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

var ans int

// removeNthFromEnd 移除链表后边第N个数
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if retrieveTree(head, n) == 0 {
		return head.Next
	}
	return head
}

// retrieveTree 后序遍历
func retrieveTree(head *ListNode, n int) int {
	if head == nil {
		return n
	}
	n = retrieveTree(head.Next, n)
	if n == 0 {
		if head.Next != nil {
			head.Next = head.Next.Next
		}
	}
	return n - 1
}
