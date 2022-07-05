package leetcode729

type MyCalendar struct {
	pairs [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{pairs: [][]int{}}
}

// 我的日程安排表
func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < len(this.pairs); i++ {
		if end > this.pairs[i][0] && start < this.pairs[i][1] {
			return false
		}
	}
	this.pairs = append(this.pairs, []int{start, end})
	return true
}
