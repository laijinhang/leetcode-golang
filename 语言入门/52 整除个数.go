/*
整除个数
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
1、2、3… …n这n(0<n<=1000000000)个数中有多少个数可以被正整数b整除。
输入
输入包含多组数据
每组数据占一行，每行给出两个正整数n、b。
输出
输出每组数据相应的结果。
样例输入
2 1
5 3
10 4
样例输出
2
1
2
 */
package main

import "fmt"

func main() {
	var n, b, num int
	var err error
	for {
		_, err = fmt.Scan(&n, &b)
		if err != nil {
			break
		}
		num = 0
		for i := 1;i <= n;i++ {
			if i % b == 0 {
				num++
			}
		}
		fmt.Println(num)
	}
}