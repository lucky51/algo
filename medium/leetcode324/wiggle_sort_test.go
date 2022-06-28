package leetcode324

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	wiggleSort(s)
	fmt.Println(s)
}
