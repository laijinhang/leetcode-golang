package main

import "math"

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := []int{math.MinInt32, 0}
	for i := 0;i < len(prices);i++ {
		dp[0] = int(math.Max(float64(dp[0]), float64(dp[1] - prices[i] - fee)))
		dp[1] = int(math.Max(float64(dp[1]), float64(dp[0] + prices[i])))
	}
	return dp[1]
}
