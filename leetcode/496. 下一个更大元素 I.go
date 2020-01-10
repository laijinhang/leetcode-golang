package main

func nextGreaterElement(findNums []int, nums []int) []int {
	res := make([]int, len(findNums))
	for i := 0;i < len(findNums);i++ {
		res[i] = -1
		for j := 0;j < len(nums);j++ {
			if findNums[i] == nums[j] {
				for ;j < len(nums);j++ {
					if findNums[i] < nums[j] {
						res[i] = nums[j]
						j = len(nums)
						break
					}
				}
			}
		}
	}
	return res
}