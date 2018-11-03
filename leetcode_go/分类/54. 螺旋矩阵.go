package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, len(matrix) * len(matrix[0]))
	n := 0
	x, y := 0, -1
	goPath := make([][]bool, len(matrix))
	for i := 0;i < len(matrix);i++ {
		goPath[i] = make([]bool, len(matrix[0]))
	}
	for n != len(matrix) * len(matrix[0]) {
		for ;y + 1 < len(matrix[0]) && !goPath[x][y+1];y++ {
			goPath[x][y+1] = true
			res[n] = matrix[x][y+1]
			n++
		}
		for ;x + 1 < len(matrix) && !goPath[x+1][y];x++ {
			goPath[x+1][y] = true
			res[n] = matrix[x+1][y]
			n++
		}
		for ;y - 1 >= 0 && !goPath[x][y-1];y-- {
			goPath[x][y-1] = true
			res[n] = matrix[x][y-1]
			n++
		}
		for ;x - 1 >= 0 && !goPath[x-1][y];x-- {
			goPath[x-1][y] = true
			res[n] = matrix[x-1][y]
			n++
		}
	}
	return res
}