package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	l := 0
	tempS := 0
	for i := 0;i < len(nums);i++ {
		tempS = 0
		for j := i;j < len(nums) && (l == 0 || j - i + 1 < l);j++ {
			tempS += nums[j]
			if tempS >= s {
				l = j - i + 1
			}
		}
	}
	return l
}

func main() {
	fmt.Println(minSubArrayLen(3, []int{1,1}))
}