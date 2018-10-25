package main

import "fmt"

var hour = [][]string{ // 1 2 4 8
	0: []string{"0"},
	1: []string{"1","2","4","8"},
	2: []string{"3","5","6", "9","10"},
	3: []string{"7", "11"},
}

var minute = [][]string {	// 1 2 4 8 16 32
	0: []string{"00"},
	1: []string{"01","02","04","08","16","32"},
	2: []string{"03","05","06","09","10","12","17","18","20","24","33","34","36","40","48"},
	3: []string{"07","11","13","14","19","21","22","25","26","28","35","37","38","41","42","44","49","50","52","56"},
	4: []string{"15","23","27","29","30","39","43","45","46","51","53","54","57","58"},
	5: []string{"31","47","55","59"},
}

func readBinaryWatch(num int) []string {
	res := []string{}
	for i := 0;i <= num && i < 4;i++ {
		if num - i <= 5 {	// 选择小时
			for _, h := range hour[i] {		// 选择分钟
				for _, m := range minute[num-i] {
					res = append(res, h + ":" + m)	// 组合，加入
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
