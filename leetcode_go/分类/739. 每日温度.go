package main

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	for i := 0;i < len(temperatures);i++ {
		j := i + 1
		for ;j < len(temperatures) && temperatures[i] >= temperatures[j];j++ {}
		if j == len(temperatures) {
			res[i] = 0
		} else {
			res[i] = j - i
		}
	}
	return res
}