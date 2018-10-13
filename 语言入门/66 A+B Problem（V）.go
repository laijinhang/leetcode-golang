/*
A+B Problem（V）
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
做了A+B Problem之后，Yougth感觉太简单了，于是他想让你求出两个数反转后相加的值。帮帮他吧
输入
有多组测试数据。每组包括两个数m和n，数据保证int范围，当m和n同时为0是表示输入结束。
输出
输出反转后相加的结果。
样例输入
1234 1234
125 117
0 0
样例输出
8642
1232
 */
package main

import (
	"fmt"
)

func main() {
	var n, m, i, j int32
	for {
		fmt.Scan(&n, &m)
		if n == 0 && m == 0 {
			break
		}
		for i = 0;n > 0;n /= 10 {
			i = i * 10 + n % 10
		}
		for j = 0;m > 0;m /= 10 {
			j = j * 10 + m % 10
		}
		fmt.Println(uint32(i) + uint32(j))
	}
}
