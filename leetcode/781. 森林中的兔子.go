package main

func numRabbits(answers []int) int {
	rabbitsM := make(map[int]int)
	ans := 0
	for i := 0;i < len(answers);i++ {
		rabbitsM[answers[i]]++
	}
	for num, count := range rabbitsM {
		if count % (num + 1) == 0 {
			ans += count
		} else {
			ans += (count / (num + 1) + 1) * (num + 1)
		}
	}
	return ans
}