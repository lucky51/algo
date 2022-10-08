package leetcodeoffer22

import . "github.com/lucky51/algo/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) (ans *ListNode) {
	var dfs func(node *ListNode)
	n := 0
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		n++
		if n == k {
			ans = node
			return
		}
	}
	dfs(head)
	return
}
