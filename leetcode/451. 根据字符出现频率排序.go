package main

import "sort"

func frequencySort(s string) string {
	chars := make(map[byte]int)
	for i := 0;i < len(s);i++ {
		chars[s[i]]++
	}
	nums := make([]int, len(chars))
	i := 0
	for _, v := range chars {
		nums[i] = v
		i++
	}
	sort.Ints(nums)
	res := ""
	for i = len(nums) - 1;i >= 0;i-- {
		for k, v := range chars {
			if nums[i] == v {
				for j := 0;j < nums[i];j++ {
					res += string(k)
				}
				delete(chars, k)
			}
		}
	}
	return res
}