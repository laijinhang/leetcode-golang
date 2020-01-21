package main

import "sort"

func lengthOfLongestSubstring(s string) int {
	maxSptLens := make([]int, len(s))
	flag := true
	for i, j := 0, 0;i < len(s);i++ {
		flag = true
		for j = i + 1;j < len(s);j++ {
			for k := i;k < j;k++ {
				if  s[k] == s[j] {
					flag = false
					break
				}
			}
			if !flag {
				break
			}
		}
		maxSptLens[i] = j - i
	}
	if len(maxSptLens) == 0 {
		return 0
	}
	sort.Ints(maxSptLens)
	return maxSptLens[len(maxSptLens)-1]
}