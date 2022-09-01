package leetcode1019

import (
	"fmt"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{7, &ListNode{5, &ListNode{1, &ListNode{9, &ListNode{2, &ListNode{5, &ListNode{Val: 1}}}}}}}}
	fmt.Println(nextLargerNodes(head))
}
