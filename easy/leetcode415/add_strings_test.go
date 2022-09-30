package leetcode415

import (
	"fmt"
	"math"
	"testing"
)

func TestStrConvertToNumber(t *testing.T) {
	numberStr := "123"
	result := 0
	for i := 0; i < len(numberStr); i++ {
		result = result*10 + int(numberStr[i]-'0')
	}
	fmt.Println(result)
}

func TestGetMaxLen(t *testing.T) {
	fmt.Println(math.MaxInt32)
}