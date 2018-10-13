/*
取石子（七）
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
Yougth和Hrdv玩一个游戏，拿出n个石子摆成一圈，Yougth和Hrdv分别从其中取石子，谁先取完者胜，每次可以从中取一个或者相邻两个，Hrdv先取，输出胜利着的名字。

输入
输入包括多组测试数据。
每组测试数据一个n，数据保证int范围内。
输出
输出胜利者的名字。
样例输入
2
3
样例输出
Hrdv
Yougth
 */

 /*
 两个不相邻的点，每次只能取1个或相邻的两个，则先取者一定输

 1）当n <= 2时，先者赢
 2）当n == 3时，后者赢
 3）当n > 3时，无论先者取1还是2，后者只要维持两段长度一样，不相邻，即后者一定赢（即奇数个时，先1，后2，先2，后1。偶数个时，先1，后1，先2，后2）
  */
package main

import "fmt"

func main() {
	var n int
	var err error
	for i := 0;i < n;i++ {
		_, err = fmt.Scan(&n)
		if err != nil {
			break
		}
		if n >= 3 {
			fmt.Println("Yougth")
		} else {
			fmt.Println("Hrdv")
		}
	}
}