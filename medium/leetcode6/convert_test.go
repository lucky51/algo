package leetcode6

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}
