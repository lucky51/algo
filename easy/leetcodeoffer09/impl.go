package leetcodeoffer09

// CQueue 用两个栈实现队列
type CQueue struct {
	s1, s2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

// AppendTail 队尾追加
func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

// 队头删除
func (this *CQueue) DeleteHead() int {
	if len(this.s2) == 0 {
		for i := len(this.s1) - 1; i >= 0; i-- {
			this.s2 = append(this.s2, this.s1[i])
		}
		this.s1 = nil
	}

	if len(this.s2) == 0 {
		return -1
	}
	v := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return v
}
