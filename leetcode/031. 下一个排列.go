package main

func nextPermutation(nums []int)  {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 1
	for ;i - 1 >= 0 && nums[i-1] >= nums[i];i-- {}  // 找出降序位置
	i--
	if i == -1 {
		sort.Ints(nums)
	} else {
		m := len(nums) - 1
		for ;m > i && nums[i] >= nums[m];m-- {}
		nums[i], nums[m] = nums[m], nums[i]
		sort.Ints(nums[i+1:])
	}
}