package main

import "strconv"

func summaryRanges(nums []int) []string {
	temps := ""
	res := []string{}
	i, j := 0, 1
	for i < len(nums) {
		temps = strconv.Itoa(nums[i])
		for j = i + 1;j < len(nums) && nums[j] - nums[j-1] == 1;j++ {
		}
		if i != j - 1 {
			temps +=  "->" + strconv.Itoa(nums[j-1])
		}
		i = j
		res = append(res, temps)
	}
	return res
}