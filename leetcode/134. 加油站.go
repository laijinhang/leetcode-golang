package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	// 找到启点
	i, b := 0, 0
	for {
		for ;i < length;i++{
			if gas[i] >= cost[i] {
				break
			}
		}
		if i == length {
			return -1
		}
		b = i
		n := 0
		for j := 0;j < length;j++ {
			n = n + gas[(j + b) % length] - cost[(j + b) % length]
			if n < 0 {
				break
			}
		}
		if n < 0 && i < length {
			i++
			continue
		}
		if n < 0 {
			return -1
		}
		return b
	}
}

func main()  {
	n := canCompleteCircuit([]int{2,3,4}, []int{3,4,3})
	fmt.Println(n)
}