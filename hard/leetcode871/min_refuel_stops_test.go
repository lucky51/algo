package leetcode871

import (
	"fmt"
	"sort"
	"testing"
)

func TestMinRefulStops(t *testing.T) {
	s := [][]int{
		{1, 0},
		{2, 1},
	}
	sort.Slice(s, func(i, j int) bool { return s[i][0] > s[j][0] })
	fmt.Println(s)

}
