/*
小学生算术
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
很多小学生在学习加法时，发现“进位”特别容易出错。你的任务是计算两个三位数在相加时需要多少次进位。你编制的程序应当可以连续处理多组数据，直到读到两个0（这是输入结束标记）。
输入
输入两个正整数m,n.(m,n,都是三位数)
输出
输出m,n,相加时需要进位多少次。
样例输入
123 456
555 555
123 594
0 0
样例输出
0
3
1
 */
package main

import "fmt"

func main() {
	var n, m int

	for {
		fmt.Scan(&n, &m)
		if n == 0 && m == 0 {
			break
		}
		a1 := [3]int{n / 100, n / 10 % 10, n % 10}
		a2 := [3]int{m / 100, m / 10 % 10, m % 10}
		count := 0
		var jw bool
		for i := 2;i >= 0;i-- {
			if jw {
				a1[i]++
				jw = false
			}
			if a1[i] + a2[i] >= 10 {
				count++
				jw = true
			}
		}
		fmt.Println(count)
	}
}