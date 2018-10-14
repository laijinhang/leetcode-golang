/*
阶乘因式分解（一）
时间限制：3000 ms  |  内存限制：65535 KB
难度：2
描述
给定两个数m,n,其中m是一个素数。

将n（0<=n<=10000）的阶乘分解质因数，求其中有多少个m。

输入
第一行是一个整数s（0<s<=100)，表示测试数据的组数
随后的s行, 每行有两个整数n，m。
输出
输出m的个数。
样例输入
2
100 5
16 2
样例输出
24
15
 */
package main

import "fmt"

func main() {
	var n, m, s, num int
	fmt.Scan(&s)
	for ;s > 0;s-- {
		fmt.Scan(&n, &m)
		num = 0
		for {
			if n % m != 0 {
				break
			}
			num += n / m
			n /= m
		}
		fmt.Println(num)
	}
}