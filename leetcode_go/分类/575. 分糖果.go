package main

func distributeCandies(candies []int) int {
	m := make(map[int]int)
	for i := 0;i < len(candies);i++ {
		m[candies[i]]++
	}
	res := 0
	if len(m) > len(candies) / 2 {
		res = len(candies) / 2
	} else {
		res = len(m)
	}
	return res
}