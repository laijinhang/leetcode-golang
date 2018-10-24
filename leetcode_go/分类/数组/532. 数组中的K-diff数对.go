package main

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	m := make(map[int]int)
	count := 0
	for i := 0;i < len(nums);i++ {
		m[nums[i]]++
	}
	for k1, _ := range m {
		if k == 0 {
			if m[k1] > 1 {
				count++
			}
		} else {
			if m[k1+k] > 0 {
				count++
			}
			if m[k1-k] > 0 {
				count++
			}
		}
		m[k1] = 0
	}
	return count
}

func main() {
	findPairs([]int{3,1,4,1,5}, 2)
}