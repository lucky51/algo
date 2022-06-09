package leetcode2

import (
	pkglist "github.com/lucky51/pkg/list"
)

func addTwoNums(l1, l2 *pkglist.Node[int]) (head *pkglist.Node[int]) {
	var tail *pkglist.Node[int]
	n1, n2, carry := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			n1 = l1.Data
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Data
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &pkglist.Node[int]{Data: sum}
			tail = head
		} else {
			tail.Next = &pkglist.Node[int]{Data: sum}
			tail = tail.Next
		}

	}
	if carry > 0 {
		tail.Next = &pkglist.Node[int]{Data: carry}
	}
	return
}
