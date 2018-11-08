package main

func search(nums []int, target int) int {
	i, j := 0, len(nums) - 1
	res := 0
	for i <= j {
		if nums[(i+j)/2] > target {
			j = (i + j) / 2 - 1
		} else if nums[(i+j)/2] < target {
			i = (i + j) / 2 + 1
		} else {
			res = (i + j) / 2
			break
		}
	}
	if i > j {
		res = -1
	}
	return res
}