package leetcode217

import (
	"fmt"
	"testing"
)

func TestQsort(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 3, 2, 1, 6, 8}
	qsort(arr)
	fmt.Printf("%v", arr)
}
