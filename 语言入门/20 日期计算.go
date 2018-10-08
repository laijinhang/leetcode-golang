/*
日期计算
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
如题，输入一个日期，格式如：2010 10 24 ，判断这一天是这一年中的第几天。
输入
第一行输入一个数N（0<N<=100）,表示有N组测试数据。后面的N行输入多组输入数据，每行的输入数据都是一个按题目要求格式输入的日期。
输出
每组输入数据的输出占一行，输出判断出的天数n
样例输入
3
2000 4 5
2001 5 4
2010 10 24
样例输出
96
124
297
 */
package main

import "fmt"

func main() {
	var N, year, month, date, n int
	fmt.Scan(&N)
	for ;N > 0;N-- {
		fmt.Scan(&year, &month, &date)
		n = 0
		for i := 1;i < month;i++ {
			switch i {
			case 1, 3, 5, 7, 8, 10, 12:
				n += 31
			case 4, 6, 9, 11:
				n += 30
			case 2:
				if (year % 4 == 0 && year % 100 != 0) || year % 100 == 0 {
					n += 29
				} else {
					n += 28
				}
			}
		}
		n += date
		fmt.Println(n)
	}
}