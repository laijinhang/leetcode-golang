/*
蛇形填数
时间限制：3000 ms  |  内存限制：65535 KB
难度：3
描述
在n*n方陈里填入1,2,...,n*n,要求填成蛇形。例如n=4时方陈为：
10 11 12 1
9 16 13 2
8 15 14 3
7 6 5 4
输入
直接输入方陈的维数，即n的值。(n<=100)
输出
输出结果是蛇形方陈。
样例输入
3
样例输出
7 8 1
6 9 2
5 4 3
 */
package main

import "fmt"

func main() {
	var matrix [101][101]int
	var n, num int
	fmt.Scan(&n)
	i, j := 0, n + 1
	for num < n * n {
		// 向下
		j--
		for i = i + 1;i <= n && matrix[i][j] == 0;i++ {
			matrix[i][j] = matrix[i-1][j] + 1
			num++
		}

		// 向左
		i--
		for j--;j > 0 && matrix[i][j] == 0;j-- {
			matrix[i][j] = matrix[i][j+1] + 1
			num++
		}

		// 向上
		j++
		for i--;i > 0 && matrix[i][j] == 0 ;i-- {
			matrix[i][j] = matrix[i+1][j] + 1
			num++
		}

		// 向右
		i++
		for j++;j <= n && matrix[i][j] == 0;j++ {
			matrix[i][j] = matrix[i][j-1] + 1
			num++
		}
	}
	for i := 1;i <= n;i++ {
		for j := 1;j <= n;j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
