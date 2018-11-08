package main

func numJewelsInStones(J string, S string) int {
	jewels := make(map[byte]bool)
	for i := 0;i < len(J);i++ {
		jewels[J[i]] = true
	}
	num := 0
	for i := 0;i < len(S);i++ {
		if jewels[S[i]] {
			num++
		}
	}
	return num
}