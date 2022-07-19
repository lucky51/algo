package leetcode731

type pair struct{ start, end int }
type MyCalendarTwo struct {
	booked, overlaps []*pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	// 先检查重叠部分，即为两次预定的部分
	for _, overlap := range this.overlaps {
		if start < overlap.end && end > overlap.start {
			return false
		}
	}
	for _, book := range this.booked {
		if start < book.end && end > book.start {
			this.overlaps = append(this.overlaps, &pair{start: max(start, book.start), end: min(end, book.end)})
		}
	}
	this.booked = append(this.booked, &pair{start: start, end: end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
