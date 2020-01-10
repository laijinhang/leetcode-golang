package main

func dominantIndex(nums []int) int {
	res := 0
	if len(nums) != 0 && len(nums) != 1 {
		max1, max2, i1 := 0, 0, 0
		if nums[0] > nums[1] {
			max1 = nums[0]
			i1 = 0
			max2 = nums[1]
		} else {
			max1 = nums[1]
			i1 = 1
			max2 = nums[0]
		}
		for i := 2;i < len(nums);i++ {
			if nums[i] >= max1 {
				max2 = max1
				i1 = i
				max1 = nums[i]
			} else if nums[i] > max2 {
				max2 = nums[i]
			}
		}
		if max1 >= 2 * max2 {
			res = i1
		} else {
			res = -1
		}
	}
	return res
}