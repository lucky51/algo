package leetcode481

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	c := []byte{2}
	fmt.Println(bytes.Repeat(c, 3))
}
