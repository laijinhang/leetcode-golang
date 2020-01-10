package main

import (
	"sort"
	"fmt"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	minutes := []int{}
	hour, minute := 0, 0
	for i := 0;i < len(timePoints);i++ {
		hour, _ = strconv.Atoi(timePoints[i][:2])
		minute, _ = strconv.Atoi(timePoints[i][3:])
		minutes = append(minutes, hour * 60 + minute)
	}
	sort.Ints(minutes)
	minutes = append(minutes, minutes[0] + 24 * 60)
	minTime := 24 * 60
	for i := 1;i < len(minutes);i++ {
		if minutes[i] - minutes[i-1] < minTime {
			minTime = minutes[i] - minutes[i-1]
		}
	}
	return minTime
}

func main() {
	as := []string{"12:00","12:01","23:00","00:00"}
	fmt.Println(findMinDifference(as))
}