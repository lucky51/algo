package leetcode707

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtTail(1)
	list.Print()
	list.Get(0)
}
