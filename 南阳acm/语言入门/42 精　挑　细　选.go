/*
精　挑　细　选
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
小王是公司的仓库管理员，一天，他接到了这样一个任务：从仓库中找出一根钢管。这听起来不算什么，但是这根钢管的要求可真是让他犯难了，要求如下：
1、 这根钢管一定要是仓库中最长的；
2、 这根钢管一定要是最长的钢管中最细的；
3、 这根钢管一定要是符合前两条的钢管中编码最大的（每根钢管都有一个互不相同的编码，越大表示生产日期越近）。
相关的资料到是有，可是，手工从几百份钢管材料中选出符合要求的那根……
要不，还是请你编写个程序来帮他解决这个问题吧。
输入
第一行是一个整数N(N<=10)表示测试数据的组数）
每组测试数据的第一行 有一个整数m(m<=1000)，表示仓库中所有钢管的数量，
之后m行，每行三个整数，分别表示一根钢管的长度（以毫米为单位）、直径（以毫米为单位）和编码（一个9位整数）。
输出
对应每组测试数据的输出只有一个9位整数，表示选出的那根钢管的编码，
每个输出占一行
样例输入
2
2
2000 30 123456789
2000 20 987654321
4
3000 50 872198442
3000 45 752498124
2000 60 765128742
3000 45 652278122
样例输出
987654321
752498124
 */
package main

import (
	"fmt"
	"sort"
)

type Pipe struct {
	length int
	diameter int
	encoding int
}

type NewPipe []Pipe

func (p NewPipe) Len() int {
	return len(p)
}

func (p NewPipe) Less(i, j int) bool {	// 判断p[i]小于p[j]，是的话返回true，否则返回false
	if p[i].length > p[j].length {	// length大的排后面
		return false
	}
	if p[i].length == p[j].length && p[i].diameter < p[j].diameter {	// diameter小的排后面
		return false
	}
	if p[i].length == p[j].length && p[i].diameter == p[j].diameter && p[i].encoding > p[j].encoding {	// encoding大的排在后面
		return false
	}
	return true
}

func (p NewPipe) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var N, m int
	p := []Pipe{}
	fmt.Scan(&N)
	for ;N > 0;N-- {
		fmt.Scan(&m)
		p = []Pipe{}
		for i := 0;i < m;i++ {
			temp := Pipe{}
			fmt.Scan(&temp.length, &temp.diameter, &temp.encoding)
			p = append(p, temp)
		}
		sort.Sort(NewPipe(p))
		fmt.Println(p[len(p)-1].encoding)
	}
}