/*
交换输出
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
输入n(n<100)个数，找出其中最小的数，将它与最前面的数交换后输出这些数。(如果这个第一个数就是最小的数，则保持原样输出，如果最小的数有相同的按照前面的交换)
输入
输入数据有多组，每组占一行，每行的开始是一个整数n，表示这个测试实例的数值的个数，跟着就是n个整数。n=0表示输入的结束，不做处理。
输出
对于每组输入数据，输出交换后的数列，每组输出占一行。
样例输入
4 2 1 3 4
5 5 4 3 2 1
0
样例输出
1 2 3 4
1 4 3 2 5
 */
package main

import "fmt"

func main() {
	var n, min int
	var arr [101]int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		min = 0
		for i := 0;i < n;i++ {
			fmt.Scan(&arr[i])
			if arr[min] > arr[i] {
				min = i
			}
		}
		arr[0], arr[min] = arr[min], arr[0]
		for i := 0;i < n;i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Println()
	}
}