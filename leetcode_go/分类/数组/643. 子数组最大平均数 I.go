package main

func findMaxAverage(nums []int, k int) float64 {
	max := -10001.0
	for i := 0;i <= len(nums) - k;i++ {
		temp := average(nums[i:i+k])
		if max < temp {
			max = temp
		}
	}
	return max
}

func average(nums []int) float64 {
	sum := 0
	for i := 0;i < len(nums);i++ {
		sum += nums[i]
	}
	return float64(sum) / float64(len(nums))
}