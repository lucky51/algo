package leetcode707

import (
	"fmt"
)

// 链表操作
// TODO: 这道题没写对，有时间在写一遍
type MyLinkedNode struct {
	Val  int
	Next *MyLinkedNode
}

type MyLinkedList struct {
	head *MyLinkedNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&MyLinkedNode{}, 0}
}

func (this *MyLinkedList) Print() {
	head := this.head
	for head != nil {
		fmt.Printf("%d=>", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

// Get 根据索引获取值
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

// AddAtHead 在链表头添加元素
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

// AddAtTail 在链表尾部添加元素
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

// AddAtIndex 在对应索引位置添加元素
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	index = max(index, 0)
	this.size++
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &MyLinkedNode{Val: val, Next: pred.Next}
	pred.Next = toAdd
}

// DeleteAtIndex 删除索引位置元素
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
