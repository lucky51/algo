package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestICVD(t *testing.T) {
	j := 0
	for i := 35; i < 60; {
		assert.Equal(t, absoluteRiskTable[j][0], absoluteRisk(0, i), "they should be equal")
		i++
		if i%5 == 0 {
			j++
		}
	}
}

var absoluteRiskTable = [][]float32{
	{1.0, 0.3},
	{1.4, 0.4},
	{1.9, 0.6},
	{2.6, 0.9},
	{3.6, 1.4},
}

func absoluteRisk(gender, age int) float32 {
	if age < 35 || age > 59 {
		return -1
	}
	return absoluteRiskTable[(age-35)/5][gender]
}
