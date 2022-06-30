package leetcode1175

const mod int = 1e9 + 7

/*
质数（Prime number），又称素数，指在大于 1 的自然数中，除了 1 和该数自身外，无法被其他自然数整除的数。
*/
// numPrimeArrangements 质数排列
func numPrimeArrangements(n int) int {
	numPrimes := 0
	// 计算n范围之内的所有质数
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			numPrimes++
		}
	}
	return factorial(numPrimes) * factorial(n-numPrimes) % mod
}

// isPrime 判断n是否为质数  [2 : sqrt(n)]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if i%2 == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i % mod
	}
	return f
}
