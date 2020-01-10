package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	res := []string{}
	tempN := append([]int{}, nums...)
	sort.Ints(tempN)
	st := make(map[int]int)
	number := 1
	for i := len(tempN) - 1;i >= 0;i-- {
		st[tempN[i]] = number
		number++
	}
	for i := 0;i < len(nums);i++ {
		if st[nums[i]] == 1 {
			res = append(res, "Gold Medal")
		} else if st[nums[i]] == 2 {
			res = append(res, "Silver Medal")
		} else if st[nums[i]] == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(st[nums[i]]))
		}
	}
	return res
}