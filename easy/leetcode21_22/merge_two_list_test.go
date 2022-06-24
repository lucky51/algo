package leetcode21_22

import (
	"fmt"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
	l3 := mergeTwoLists(l1, l2)
	for temp := l3; temp != nil; temp = temp.Next {
		fmt.Println(temp.Val)
	}

}
