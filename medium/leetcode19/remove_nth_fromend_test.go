package leetcode19

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// h := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	h1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	newHead := removeNthFromEnd(h1, 2)
	fmt.Println("print list")
	for temp := newHead; temp != nil; temp = temp.Next {
		fmt.Println(temp.Val)
	}
}
