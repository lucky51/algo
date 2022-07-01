package leetcode160

import "github.com/lucky51/algo/common"

// getIntersectionNode 相交链表 ，利用哈希字典表记录访问过的节点
func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	seen := map[*common.ListNode]struct{}{}
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := seen[headA]; ok {
				return headA
			}
			seen[headA] = struct{}{}

			headA = headA.Next
		}
		if headB != nil {
			if _, ok := seen[headB]; ok {
				return headB
			}
			seen[headB] = struct{}{}
			headB = headB.Next
		}
	}
	return nil
}
