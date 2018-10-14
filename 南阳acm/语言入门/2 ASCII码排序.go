package main

import "fmt"

/*
ASCII码排序
时间限制：3000 ms  |  内存限制：65535 KB
难度：2
描述
输入三个字符（可以重复）后，按各字符的ASCII码从小到大的顺序输出这三个字符。
输入
第一行输入一个数N,表示有N组测试数据。后面的N行输入多组数据，每组输入数据都是占一行，有三个字符组成，之间无空格。
输出
对于每组输入数据，输出一行，字符中间用一个空格分开。
样例输入
2
qwe
asd
样例输出
e q w
a d s
 */

 func main() {
 	var n int
 	var str string

 	fmt.Scan(&n)
 	for ;n > 0;n-- {
 		fmt.Scan(&str)
 		b := []byte(str)
 		for i := 0;i < 3;i++ {
 			t := minChar(b[i:]) + i
			b[i], b[t] = b[t], b[i]
		}
		fmt.Println(string(b))
	}
 }

 func minChar(b []byte) int {
	 var xb int
 	length := len(b)
 	for i := 0;i < length;i++ {
 		if b[xb] > b[i] {
 			xb = i
		}
	}
	 return xb
 }