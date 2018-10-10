/*
分数加减法
时间限制：1000 ms  |  内存限制：65535 KB
难度：2
描述
编写一个C程序，实现两个分数的加减法
输入
输入包含多行数据
每行数据是一个字符串，格式是"a/boc/d"。
其中a, b, c, d是一个0-9的整数。o是运算符"+"或者"-"。

数据以EOF结束
输入数据保证合法
输出
对于输入数据的每一行输出两个分数的运算结果。
注意结果应符合书写习惯，没有多余的符号、分子、分母，并且化简至最简分数
样例输入
1/8+3/8
1/4-1/2
1/3-1/3
样例输出
1/2
-1/4
0
 */
package main

import "fmt"

func main() {
	var a, b, c, d int
	var negative bool
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		negative = false
		a = int(str[0] - '0')
		b = int(str[2] - '0')
		c = int(str[4] - '0')
		d = int(str[6] - '0')
		if str[3] == '-' {
			a = a * d - c * b
			if a < 0 {
				a = -a
				negative = true
			}
		} else {
			a = a * d + c * b
		}
		b *= d
		if a == 0 {
			fmt.Printf("0\n")
			continue
		}
		for i := a;i > 1;i-- {
			if a % i == 0 && b % i == 0 {
				a /= i
				b /= i
			}
		}
		if negative {
			a = -a
		}

		fmt.Printf("%d/%d\n", a, b)
	}
}