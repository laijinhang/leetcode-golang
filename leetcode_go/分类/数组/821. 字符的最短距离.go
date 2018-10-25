package main

import (
	"fmt"
	"math"
)

func shortestToChar(S string, C byte) []int {
	cIndex := []int{}
	for i := 0;i < len(S);i++ {
		if S[i] == C {
			cIndex = append(cIndex, i)
		}
	}
	fmt.Println(cIndex)
	res := []int{}
	for i := 0;i < len(S);i++ {
		min := len(S)
		for j := 0;j < len(cIndex);j++ {
			if min == 0 {
				break
			}
			if min > int(math.Abs(float64(i - cIndex[j]))) {
				min = int(math.Abs(float64(i - cIndex[j])))
			}
		}
		res = append(res, min)
	}
	return res
}