/*
背包问题
时间限制：3000 ms  |  内存限制：65535 KB
难度：3
描述
现在有很多物品（它们是可以分割的），我们知道它们每个物品的单位重量的价值v和重量w（1<=v,w<=10）；如果给你一个背包它能容纳的重量为m（10<=m<=20）,你所要做的就是把物品装到背包里，使背包里的物品的价值总和最大。
输入
第一行输入一个正整数n（1<=n<=5）,表示有n组测试数据；
随后有n测试数据，每组测试数据的第一行有两个正整数s，m（1<=s<=10）;s表示有s个物品。接下来的s行每行有两个正整数v，w。
输出
输出每组测试数据中背包内的物品的价值和，每次输出占一行。
样例输入
1
3 15
5 10
2 8
3 9
样例输出
65
 */
package main

import "fmt"

type goods struct {
	v, w int
}

var maxV int	// 最大价值
var maxW int	// 背包可允许的最大重

func bb(gs []goods, n, v, w int) {
	if v > maxV {
		maxV = v
	}
	if n >= len(gs) {	// 退出
		return
	}
	// 不装进书包
	bb(gs, n+1, v, w)
	// 装进一件物品到书包
	if gs[n].w > 0 && w + 1 <= maxW {
		/* 深度拷贝切片数据 */
		tgs := make([]goods, 0, len(gs))
		for _, g := range gs {
			tgs = append(tgs, g)
		}
		tgs[n].w--	// 物品减一
		bb(tgs, n, v+gs[n].v, w+1)
	}
}

func main() {
	var n, s int
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		fmt.Scan(&s, &maxW)
		maxV = 0
		tempg := goods{}
		gs := make([]goods, 0, s)
		for i := 0;i < s;i++ {
			fmt.Scan(&tempg.v, &tempg.w)
			gs = append(gs, tempg)
		}
		bb(gs, 0, 0, 0)
		fmt.Println(maxV)
	}
}