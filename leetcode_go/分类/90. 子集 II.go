package main

import (
	"sort"
	"fmt"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	swd(nums, []int{}, 0, &res)
	return res
}

func swd(nums, r []int, i int, res *[][]int) {
	if i >= len(nums) {
		*res = append(*res, r)
		return
	}
	j := i + 1
	for ;j < len(nums);j++ {
		if nums[i] != nums[j] {
			break
		}
	}
	swd(nums, append([]int{}, r...), j, res)    // 不加
	swd(nums, append(r, nums[i]), i + 1, res)   // 加
}

func main() {
	fmt.Println(subsetsWithDup([]int{2,1,2,1,3}))
}