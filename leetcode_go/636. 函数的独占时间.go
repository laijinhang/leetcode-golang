package main

import (
	"strings"
	"strconv"
	"fmt"
)

type funcTask struct {
	b int   // 开始时间
	fId int // 函数id
}

func exclusiveTime(n int, logs []string) []int {
	funcStack := make([]funcTask, len(logs))     // 栈
	funcSpanTime := make([]int, len(logs))       // 调用其中栈所消耗时间
	res := make([]int, n)
	fsPoint := 0
	for i := 0;i < len(logs);i++ {
		logData := strings.Split(logs[i], ":")
		fId, _ := strconv.Atoi(logData[0])
		t, _ := strconv.Atoi(logData[2])
		if logData[1] == "start" {
			funcStack[fsPoint] = funcTask{b:t,fId:fId}
			fsPoint++
		} else {
			fsPoint--
			tn := t - funcStack[fsPoint].b + 1 - funcSpanTime[fsPoint]
			res[fId] += tn
			funcSpanTime[fsPoint] = 0
			for i := 0;i < fsPoint;i++ {
				funcSpanTime[i] += tn
			}
		}
	}
	return res
}

func main() {
	fmt.Println(exclusiveTime(2, []string{
		"0:start:0",
		"1:start:2",
		"1:end:3",
		"0:start:4",
		"0:end:5",
		"0:end:6"}))
}