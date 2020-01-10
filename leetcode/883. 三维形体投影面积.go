package main

func projectionArea(grid [][]int) int {
	s1, s2, s3 := 0, 0, 0
	for i := 0;i < len(grid);i++ {
		max := grid[i][0]
		for j := 0;j < len(grid[0]);j++ {
			if grid[i][j] != 0 {
				s1++
			}
			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
		s2 += max
	}
	for i := 0;i < len(grid[0]);i++ {
		max := grid[0][i]
		for j := 1;j < len(grid);j++ {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		s3 += max
	}
	return s1 + s2 + s3
}