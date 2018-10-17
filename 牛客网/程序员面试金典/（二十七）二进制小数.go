/*
题目描述
有一个介于0和1之间的实数，类型为double，返回它的二进制表示。如果该数字无法精确地用32位以内的二进制表示，返回“Error”。

给定一个double num，表示0到1的实数，请返回一个string，代表该数的二进制表示或者“Error”。

测试样例：
0.625
返回：0.101

分析：
小数左移一位，如果值大于1，说明小数点后一位是1，如果小于1，说明小数点后一位是0

 */
package main

import "fmt"

func printBin(num float64) string {
	s := "0."
	for num > 0 {
		num *= 2
		if num >= 1 {
			s += "1"
			num--
		} else {
			s += "0"
		}
		if len(s) > 32 {
			return "Error"
		}
	}
	return s
}

func main() {
	b := printBin(0.625)
	fmt.Println(b)
}