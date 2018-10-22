package main

func twoSum(numbers []int, target int) []int {
	index1, index2 := 0, 0
	for i := 0;i < len(numbers);i++ {
		for j := i + 1;j < len(numbers);j++ {
			if numbers[i] + numbers[j] == target {
				index1 = i + 1
				index2 = j + 1
				i = len(numbers)
				j = len(numbers)
			} else if numbers[i] + numbers[j] > target {
				break
			}
		}
	}
	return []int{index1,index2}
}