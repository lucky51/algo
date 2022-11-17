package leetcode792

import (
	"fmt"
	"sort"
	"testing"
)

func TestSearchInts(t *testing.T) {
	var arr = []int{2, 3, 4, 5, 6, 8, 10}
	fmt.Println(sort.SearchInts(arr, 7))
	var arr1 = []int{2}
	fmt.Println(sort.SearchInts(arr1, 3))
}

func TestNumMatchingSubSeq(t *testing.T) {
	numMatchingSubseq("abcde", []string{"ace", "cde", "de"})
}
