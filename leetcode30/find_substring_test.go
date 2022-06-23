package leetcode30

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

func TestMap(t *testing.T) {
	m := map[string]int{}
	m["123"]++
	fmt.Println(m["123"])
}
