package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	left, right := 0, len(matrix) * len(matrix[0]) - 1
	for left <= right {
		if matrix[((left + right) / 2) / len(matrix[0])][((left + right) / 2) % len(matrix[0])] == target {
			return true
		} else if matrix[(left + right) / (2 * len(matrix[0]))][((left + right) / 2) % len(matrix[0])] < target {
			left = (left + right) / 2 + 1
		} else {
			right = (left + right) / 2 - 1
		}
	}
	return false
}