package main

import (
	"sort"
)

func removeDuplicates(nums []int) int {
	sort.Ints(nums)
	nLen := len(nums)
	n, i, j := 1, 0, 0
	for ;i < nLen - 1;i++ {
		for j = i + 1;j < nLen && nums[i] == nums[j];j++ {}
		if j == nLen {
			break
		}
		nums[n] = nums[j]
		if i != j - 1 {
			i = j - 1
		}
		n++
	}
	return n
}

func main()  {
	removeDuplicates([]int{1,1,1,1,1})
}