package main

func maxProfit(prices []int) int {
	max := 0
	for i := 0;i < len(prices);i++ {
		for j := i + 1;j < len(prices);j++ {
			if prices[j] - prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}