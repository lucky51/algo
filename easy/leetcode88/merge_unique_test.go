package leetcode88

import (
	"fmt"
	"testing"
)

func TestMergeUnique(t *testing.T) {
	a, b := []int{9, 22, 35, 72, 100, 200}, []int{1, 9, 22, 72, 200}
	fmt.Println(mergeUnique1(a, b))
}
