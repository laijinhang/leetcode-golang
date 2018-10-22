package main

func majorityElement(nums []int) int {
	nLen := len(nums)
	num := 0
	res := 0
	for i := 0;i < nLen / 2 + 1;i++ {
		num = 1
		for j := i + 1;j < nLen;j++ {
			if nums[j] == nums[i] {
				num++
			}
		}
		if num > nLen / 2 {
			res = i
			break
		}
	}
	return nums[res]
}