/*
喷水装置（二）
时间限制：3000 ms  |  内存限制：65535 KB
难度：4
描述
有一块草坪，横向长w,纵向长为h,在它的橫向中心线上不同位置处装有n(n<=10000)个点状的喷水装置，每个喷水装置i喷水的效果是让以它为中心半径为Ri的圆都被润湿。请在给出的喷水装置中选择尽量少的喷水装置，把整个草坪全部润湿。
输入
第一行输入一个正整数N表示共有n次测试数据。
每一组测试数据的第一行有三个整数n,w,h，n表示共有n个喷水装置，w表示草坪的横向长度，h表示草坪的纵向长度。
随后的n行，都有两个整数xi和ri,xi表示第i个喷水装置的的横坐标（最左边为0），ri表示该喷水装置能覆盖的圆的半径。
输出
每组测试数据输出一个正整数，表示共需要多少个喷水装置，每个输出单独占一行。
如果不存在一种能够把整个草坪湿润的方案，请输出0。
样例输入
2
2 8 6
1 1
4 5
2 10 6
4 5
6 5
样例输出
1
2
 */
package main

import (
	"fmt"
	"math"
)

type zb struct {
	b float64
	e float64
}

func main() {
	var N, n, w, h int
	var x, r float64
	fmt.Scan(&N)
	for i := 0;i < N;i++ {
		fmt.Scan(&n, &w, &h)
		zbs := []zb{}
		for i := 0;i < n;i++ {
			fmt.Scan(&x, &r)
			if int(2 * r * 1000) <= h * 1000 {
				continue
			}
			l := math.Sqrt(r * r - float64(h) * float64(h) / 4.0)
			zbs = append(zbs, zb{x - l, x + l})
		}
		num := 0
		for tempW := float64(0);int(tempW * 1000) < w * 1000;{
			// 找出当前位置出发的最长路径
			tempZb := zb{}
			maxLen := float64(0)
			for _, v := range zbs {
				if v.b <= tempW && v.e - tempW > maxLen {
					maxLen = v.e - tempW
					tempZb = v
				}
			}
			if tempZb.e == tempZb.b && tempZb.b == float64(0) {
				num = 0
				break
			}
			tempW = tempZb.e
			num++
		}
		fmt.Println(num)
	}
}
