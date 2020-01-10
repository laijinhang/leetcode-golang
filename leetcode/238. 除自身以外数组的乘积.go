package main

import "fmt"

func productExceptSelf(nums []int) []int {
	fs, ps := make([]int, len(nums)), make([]int, len(nums))
	f, p := 1, 1
	for i := 0;i < len(nums);i++ {
		f *= nums[i]
		fs[i] = f
		p *= nums[len(nums) -i-1]
		ps[i] = p
	}
	fmt.Println(fs)
	fmt.Println(ps)
	t := make([]int, len(nums))
	for i := 0;i < len(nums);i++ {
		if i == 0 {
			t[i] = ps[i+1]
		} else if i == len(nums) - 1 {
			t[i] = fs[i-1]
		} else {
			t[i] = fs[i-1] * ps[i+1]
		}
	}
	return t
}