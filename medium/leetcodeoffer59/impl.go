package leetcodeoffer59

import "math"

type MaxQueue struct {
	queue []int
}

func Constructor() MaxQueue {
	return MaxQueue{queue: []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	maxValue := math.MinInt32
	for i := 0; i < len(this.queue); i++ {
		if this.queue[i] > maxValue {
			maxValue = this.queue[i]
		}
	}
	return maxValue
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	v := this.queue[0]
	this.queue = this.queue[1:]
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
