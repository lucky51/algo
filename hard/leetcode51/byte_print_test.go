package leetcode51

import (
	"fmt"
	"testing"
)

func TestPrintBytes(t *testing.T) {
	b := []byte{'Q', '.'}
	fmt.Println(string(b))
	fmt.Printf("%s", b)
}
