package leetcode20

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("[[]]()"))
	fmt.Println(isValid("[[]]()}"))
}
