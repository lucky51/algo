package leetcode901

// StockSpanner 股票价格跨度
type StockSpanner struct {
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	ans := 1
	for i := len(this.stack) - 1; i >= 0 && this.stack[i] <= price; i-- {
		ans++
	}
	this.stack = append(this.stack, price)
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
