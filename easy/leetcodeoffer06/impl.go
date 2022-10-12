package leetcodeoffer06

import . "github.com/lucky51/algo/common"

// reversePrint 剑指 Offer 06. 从尾到头打印链表
func reversePrint(head *ListNode) (ans []int) {
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		ans = append(ans, node.Val)
	}
	dfs(head)
	return
}
