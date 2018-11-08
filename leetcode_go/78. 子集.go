package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	var subs func(nums, sub []int, i int)
	subs = func (nums, sub []int, i int) {
		if i == len(nums) {
		res = append(res, sub)
		return
		}
		subs(nums, sub, i + 1)
		subs(nums, append(sub, nums[i]), i + 1)
	}
	subs(nums, []int{}, 0)
	fmt.Println(res)
	return res
}


func main()  {
	fmt.Println(subsets([]int{9,0,3,5,7}))
}