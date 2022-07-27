package leetcode592

import (
	"fmt"
	"unicode"
)

// fractionAddition 分数加减运算
func fractionAddition(expression string) string {
	// 几个词 ：denominator 分母 ，numerator 分子 ， greatest common divisor(gcd) 最大公约数
	// 一个基本的数学公式   numerator1 /denominator1  +numerator2 /denominator2 = numerator1 *denominator2 + numerator2 * denominator1 /  denominator1 * denominator2
	// 提取分子和分母按照公式进行计算
	numerator, denominator := 0, 1 // 分子，分母
	for i, n := 0, len(expression); i < n; {
		// 读取分子
		numerator1, sign := 0, 1
		if expression[i] == '+' || expression[i] == '-' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}

		for i < n && unicode.IsDigit(rune(expression[i])) {
			numerator1 = numerator1*10 + int(expression[i]-'0')
			i++
		}
		numerator1 *= sign
		// 跳过 /
		i++
		// 读取分母
		denominator1 := 0
		for i < n && unicode.IsDigit(rune(expression[i])) {
			denominator1 = denominator1*10 + int(expression[i]-'0')
			i++
		}
		// 计算分子
		numerator = numerator*denominator1 + numerator1*denominator
		// 计算分母
		denominator *= denominator1
	}
	if numerator == 0 {
		return "0/1"
	}
	g := gcd(abs(numerator), denominator)
	return fmt.Sprintf("%d/%d", numerator/g, denominator/g)
}

// 求最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
