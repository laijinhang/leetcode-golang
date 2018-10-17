/*
题目描述
有一副由NxN矩阵表示的图像，这里每个像素用一个int表示，请编写一个算法，在不占用额外内存空间的情况下(即不使用缓存矩阵)，将图像顺时针旋转90度。

给定一个NxN的矩阵，和矩阵的阶数N,请返回旋转后的NxN矩阵,保证N小于等于500，图像元素小于等于256。

测试样例：
[[1,2,3],[4,5,6],[7,8,9]],3
返回：[[7,4,1],[8,5,2],[9,6,3]]

解析：顺时针翻转90度相当于先上下 翻转，再沿主对角线翻转，由此分为两个步骤：
1）上下翻转
  （i, j) -> (n-i-1, j)
2）沿主对角线翻转（从左上角到右下角）
    (i, j) -> (j, i)
思路来源：https://zhuanlan.zhihu.com/p/26548146
 */
package main

import "fmt"

func transformImage(mat [500][500]int, n int) [500][500]int {
	var i, j int
	for ;i < n / 2 && i != n - i - 1;i++ {
		// 上下翻转
		for j = 0;j < n;j++ {
			mat[i][j], mat[n-i-1][j] = mat[n-i-1][j], mat[i][j]
		}
	}
	// 沿主对角线翻转
	for i = 0;i < n;i++ {
		for j = i + 1;j < n;j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	return mat
}

func main()  {
	mat := [500][500]int{{1,2,3},{4,5,6},{7,8,9}}
	mat = transformImage(mat, 3)

	for i := 0;i < 3;i++ {
		for j := 0;j < 3;j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}