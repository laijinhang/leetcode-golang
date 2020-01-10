package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	nM := make(map[int]int)
	for i := 0;i < len(nums);i++ {
		nM[nums[i]]++
	}
	counts := make([]int, len(nums))
	i := 0
	for _, v := range nM {
		counts[i] = v
		i++
	}
	res := []int{}
	sort.Ints(counts)
	for i := len(counts) - 1;i >= len(counts) - k;i-- {
		for k, v := range nM {
			if v == counts[i] {
				res = append(res, k)
				delete(nM, k)
			}
		}
	}
	return res
}