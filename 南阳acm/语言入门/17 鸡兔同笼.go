/*
鸡兔同笼
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
已知鸡和兔的总数量为n,总腿数为m。输入n和m,依次输出鸡和兔的数目，如果无解，则输出“No answer”(不要引号)。
输入
第一行输入一个数据a,代表接下来共有几组数据，在接下来的(a<10)
a行里，每行都有一个n和m.(0<m,n<100)
输出
输出鸡兔的个数，或者No answer
样例输入
2
14 32
10 16
样例输出
12 2
No answer
 */
package main

import "fmt"

func main() {
	var a, n, m, x, y int
	fmt.Scan(&a)
	for ;a > 0;a-- {
		fmt.Scan(&n, &m)
		for x = 0;x <= n;x++ {
			y = n - x
			if 2 * x + 4 * y == m {
				break
			}
		}
		if x <= n {
			fmt.Println(x, y)
		} else {
			fmt.Println("No answer")
		}
	}
}
