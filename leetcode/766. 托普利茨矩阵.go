package main

/*
判断邻接对角线上是否相等，如果不相等，然后false
 */
func isToeplitzMatrix(matrix [][]int) bool {
	rLen := len(matrix)
	cLen := len(matrix[0])
	for i := 0;i < rLen - 1;i++ {
		for j := 0;j < cLen - 1;j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}