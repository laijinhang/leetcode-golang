/*
计算球体积
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
根据输入的半径值，计算球的体积。
输入
输入数据有多组，每组占一行，每行包括一个实数，表示球的半径。（0<R<100）
输出
输出对应的球的体积，对于每组输入数据，输出一行，计算结果四舍五入为整数
Hint:PI=3.1415926
样例输入
1
1.5
样例输出
4
14
 */
package main

import "fmt"

const PI = 3.1415926

func main() {
	var R float64
	for {
		_, err := fmt.Scan(&R)
		if err != nil {
			break
		}
		fmt.Println(int(4 / 3.0 * PI * float64(R * R * R)))
	}
}