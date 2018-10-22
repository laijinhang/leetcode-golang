package main

type group struct {
	i1, i2 int
}

func largeGroupPositions(S string) [][]int {
	res := [][]int{}
	temp := []int{0,0}
	for i := 0;i < len(S) - 1;i++ {
		if S[i] == S[i+1] {
			temp[1]++
		} else {
			if temp[1] - temp[0] >= 2 {
				res = append(res, temp)
			}
			temp = []int{i+1,i+1}
		}
	}
	if temp[1] - temp[0] >= 2 {
		res = append(res, temp)
	}
	return res
}