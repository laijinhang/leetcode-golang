/*
开灯问题
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
有n盏灯，编号为1~n，第1个人把所有灯打开，第2个人按下所有编号为2 的倍数的开关（这些灯将被关掉），第3 个人按下所有编号为3的倍数的开关（其中关掉的灯将被打开，开着的灯将被关闭），依此类推。一共有k个人，问最后有哪些灯开着？输入：n和k，输出开着的灯编号。k≤n≤1000

输入
输入一组数据：n和k
输出
输出开着的灯编号
样例输入
7 3
样例输出
1 5 6 7
 */
 package main

import "fmt"

func main() {
	var n, k int
	var d [1001]bool
	fmt.Scan(&n, &k)
	for i := 1;i <= k;i++ {
		for j := i;j <= n;j += i {
			d[j] = !d[j]
		}
	}
	for i := 1;i <= n;i++ {
		if d[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}