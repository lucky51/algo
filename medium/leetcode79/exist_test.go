package leetcode79

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	fmt.Println(exist([][]byte{
		{'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A'}}, "AAAAAAAAAAAAAAB"))
}
