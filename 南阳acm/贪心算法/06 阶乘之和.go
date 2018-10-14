/*
阶乘之和
时间限制：3000 ms  |  内存限制：65535 KB
难度：3
描述
给你一个非负数整数n，判断n是不是一些数（这些数不允许重复使用，且为正数）的阶乘之和，如9=1！+2!+3!，如果是，则输出Yes，否则输出No；

输入
第一行有一个整数0<m<100,表示有m组测试数据；
每组测试数据有一个正整数n<1000000;
输出
如果符合条件，输出Yes，否则输出No;
样例输入
2
9
10
样例输出
Yes
No
 */
package main

import "fmt"

func main() {
	s := make(map[int]bool)
	for n, x, i := 0, 1, 1;n <= 1000000;i++ {
		x *= i
		n += x
		fmt.Println(n)
		s[n] = true
	}
	var m, n int
	fmt.Scan(&m)
	for i := 0;i < m;i++ {
		fmt.Scan(&n)
		if s[n] == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}