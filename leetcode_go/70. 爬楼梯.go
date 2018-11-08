package main

func climbStairs(n int) int {
	ways := []int{1,1}
	for i := 2;i <= n;i++ {
		ways[i % 2] = ways[0] + ways[1]
	}
	return ways[n % 2]
}
