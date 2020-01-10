package main

import "strconv"

func calPoints(ops []string) int {
	points := 0
	calP := make([]int, len(ops))
	ci := 0
	for i := 0;i < len(ops);i++ {
		switch ops[i] {
		case "C":
			points -= calP[ci]
			ci--
		case "D":
			calP[ci+1] = calP[ci] * 2
			ci++
			points += calP[ci]
		case "+":
			calP[ci+1] = calP[ci-1] + calP[ci]
			ci++
			points += calP[ci]
		default:
			ci++
			calP[ci], _ = strconv.Atoi(ops[i])
			points += calP[ci]
		}
	}
	return points
}