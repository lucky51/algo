package leetcode19

import "fmt"

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	if retrieveTree(head, n) == 0 {
		return head.Next
	}

	return head
}
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
	fmt.Println("node:", head.Val, n)
	return n - 1
}
