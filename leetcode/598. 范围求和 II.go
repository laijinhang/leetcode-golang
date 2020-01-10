package main

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	r, c := 40001, 40001
	for i := 0;i < len(ops);i++ {
		if r > ops[i][0] {
			r = ops[i][0]
		}
		if c > ops[i][1] {
			c = ops[i][1]
		}
	}
	return r * c
}