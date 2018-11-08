package main

import "fmt"

/*
分析：其实就是将其分成最长递增子序列
 */
func maxProfit(prices []int) int {
	tempP := [][]int{}
	for i := 0;i < len(prices);i++ {
		j := i + 1
		for ;j < len(prices) && prices[j] >= prices[j - 1];j++ {}
		j--
		if j > i {
			tempP = append(tempP, prices[i:j+1])
			i = j
		}
	}
	if len(tempP) == 0 {
		return 0
	}
	num := 0
	for _, p := range tempP {
		num += (p[len(p)-1] - p[0])
	}
	return num
}

func main() {
	n := maxProfit([]int{7, 1, 4, 3, 1})
	fmt.Println(n)
}
