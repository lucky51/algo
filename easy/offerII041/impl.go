package offerii041

// 滑动窗口的平均值
type MovingAverage struct {
	windows []int
	size    int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:    size,
		windows: []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.windows = append(this.windows, val)
	sum := 0
	for i, cnt := len(this.windows)-1, this.size-1; i >= 0 && cnt >= 0; i, cnt = i-1, cnt-1 {
		sum += this.windows[i]
	}
	if len(this.windows) < this.size {
		return float64(sum) / float64(len(this.windows))
	}
	return float64(sum) / float64(this.size)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
