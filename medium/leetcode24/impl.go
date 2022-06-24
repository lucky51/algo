package leetcode24

import "github.com/lucky51/algo/common"

// swapPairs 两两交换链表中的节点
func swapPairs(head *common.ListNode) *common.ListNode {
	//base case 如果节点为空或者是只有一个单节点，就返回自己
	if head == nil || head.Next == nil {
		return head
	}
	one := head
	two := one.Next
	three := two.Next
	// 新的头结点为第二个节点
	two.Next = one
	// 第三个节点传入方法得到下一次交换的结果，链接第二个节点，即新链表的头结点
	one.Next = swapPairs(three)
	return two
}
