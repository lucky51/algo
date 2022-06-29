package leetcode83

import (
	"github.com/lucky51/algo/common"
)

// deleteDuplicates 删除排序链表中的重复元素
func deleteDuplicates(head *common.ListNode) *common.ListNode {
	for temp := head; temp != nil && temp.Next != nil; {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head
}
