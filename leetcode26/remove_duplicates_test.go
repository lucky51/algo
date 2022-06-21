package leetcode26

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(arr)
	fmt.Println(arr)
}
