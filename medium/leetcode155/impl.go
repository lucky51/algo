package leetcode155

import "math"

type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minstack: []int{math.MaxInt32}, // 添加一个辅助栈，用于记录每一个push操作对应的最小值
	}
}

func (this *MinStack) Push(val int) {
	minTop := this.minstack[len(this.minstack)-1]
	this.minstack = append(this.minstack, min(val, minTop))
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() (min int) {
	return this.minstack[len(this.minstack)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
