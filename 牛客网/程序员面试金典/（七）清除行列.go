/*
题目描述
请编写一个算法，若N阶方阵中某个元素为0，则将其所在的行与列清零。

给定一个N阶方阵int[][](C++中为vector<vector><int>>)mat和矩阵的阶数n，请返回完成操作后的int[][]方阵(C++中为vector<vector><int>>)，保证n小于等于300，矩阵中的元素为int范围内。</int></vector></int></vector>

测试样例：
[[1,2,3],[0,1,2],[0,0,1]]
返回：[[0,0,3],[0,0,0],[0,0,0]]
 */
package main

import "fmt"

func clearZero(mat [500][500]int, n int) [500][500]int {
	var m [500][500]bool
	for i := 0;i < n;i++ {
		for j := 0;j < n;j++ {
			if mat[i][j] == 0 {
				for k := 0;k < n;k++ {
					m[i][k] = true
					m[k][j] = true
				}
			}
		}
	}
	for i := 0;i < n;i++ {
		for j := 0; j < n; j++ {
			if m[i][j] {
				mat[i][j] = 0
			}
		}
	}

	return mat
}

func main() {
	mat := clearZero([500][500]int{{1,2,3},{0,1,2},{0,0,1}}, 3)
	for i := 0;i < 3;i++ {
		for j := 0;j < 3;j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}