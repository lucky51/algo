package leetcode141

import "github.com/lucky51/algo/common"

// hasCycle环形链表 快慢指针
func hasCycle(head *common.ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// hasCycle1 环形链表用hash记录访问过的节点
func hasCycle1(head *common.ListNode) bool {
	// 记录
	seen := map[*common.ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		head = head.Next
	}
	return false
}
