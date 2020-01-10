package main

func setZeroes(matrix [][]int)  {
	is, js := make(map[int]bool), make(map[int]bool)
	for i := 0;i < len(matrix);i++ {
		for j := 0;j < len(matrix[i]);j++ {
			if matrix[i][j] == 0 {
				is[i] = true
				js[i] = true
			}
		}
	}
	for i := 0;i < len(matrix);i++ {	// 行清空
		for j := 0;j < len(matrix[i]) && is[i];j++ {
			matrix[i][j] = 0
		}
	}
	for i := 0;i < len(matrix[0]);i++ {
		if js[i] {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
}