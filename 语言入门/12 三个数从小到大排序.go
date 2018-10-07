/*
三个数从小到大排序
时间限制：3000 ms  |  内存限制：65535 KB
难度：0
描述
现在要写一个程序，实现给三个数排序的功能

输入
输入三个正整数
输出
给输入的三个正整数排序
样例输入
20 7 33
样例输出
7 20 33
 */
package main

import "fmt"

func main() {
	var i, j, k int
	fmt.Scan(&i, &j, &k)
	if i >= j && j >= k {
		fmt.Println(k, j, i)
	} else if i >= k && k >= j {
		fmt.Println(j, k, i)
	} else if j >= i && i >= k {
		fmt.Println(k, i, j)
	} else if j >= k && k >= i {
		fmt.Println(i, k, j)
	} else if k >= i && i >= j {
		fmt.Println(j, i, k)
	} else if k >= j && j >= i {
		fmt.Println(i, j, k)
	}
}