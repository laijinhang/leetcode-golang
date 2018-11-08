package main

import "fmt"

func countPrimes(n int) int {
	count := 0
	prime := make([]bool, n+1)
	for i := 2;i <= n;i++ {
		prime[i] = true
	}
	for i := 2;i * i <= n;i++ {
		if prime[i] {
			for j := i + i; j <= n; j += i {
				prime[j] = false
			}
		}
	}
	for i := 2;i < n;i++ {
		if prime[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10000))
}