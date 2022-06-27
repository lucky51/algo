package leetcode522

import (
	"fmt"
	"testing"
)

func TestFindLUSLength(t *testing.T) {
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
}

func TestContinueOuter(t *testing.T) {
	arr := []int{1, 2, 3, 4}
next:
	for i, s := range arr {
		for j, t := range arr {
			fmt.Println(i, s, j, t)
			if i == 1 {
				fmt.Println("continue:::")
				continue next
			}
		}
	}
}
