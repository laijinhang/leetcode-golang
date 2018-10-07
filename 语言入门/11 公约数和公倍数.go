/*
讨论区
公约数和公倍数
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
小明被一个问题给难住了，现在需要你帮帮忙。问题是：给出两个正整数，求出它们的最大公约数和最小公倍数。
输入
第一行输入一个整数n（0<n<=10000)，表示有n组测试数据;
随后的n行输入两个整数i,j（0<i,j<=32767)。
输出
输出每组测试数据的最大公约数和最小公倍数
样例输入
3
6 6
12 11
33 22
样例输出
6 6
1 132
11 66
 */
package main

import "fmt"

func main() {
	var n, i, j, bcs, ys, cs int
	fmt.Scan(&n)
	for ;n > 0;n-- {
		fmt.Scan(&i, &j)
		if i > j {
			bcs = i
			cs = j
		} else {
			bcs = j
			cs = i
		}
		for ;cs != 0; {
			ys = bcs % cs
			bcs = cs
			cs = ys
		}
		fmt.Println(bcs, i * j / bcs)
	}
}
