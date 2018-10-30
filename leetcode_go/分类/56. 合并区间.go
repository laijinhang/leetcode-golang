package main

import (
	"sort"
	"fmt"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type Interval struct {
	Start int
	End   int
}

type IntervalS []Interval

func (this IntervalS) Len() int {
	return len(this)
}

func (this IntervalS) Less(i, j int) bool {
	if this[i].Start > this[j].Start || (this[i].Start == this[j].Start && this[i].End > this[j].End) {
		return false
	}
	return true
}

func (this IntervalS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func merge(intervals []Interval) []Interval {
	// 1）重排
	sort.Sort(IntervalS(intervals))
	res := []Interval{}
	// 2）合并和去重
	intervalM := make(map[Interval]bool)
	for i := 0;i < len(intervals);i++ {
		for j := 0;j < len(intervals);j++ {
			if i == j {
				continue
			}
			if intervals[i].End < intervals[j].Start {
				break
			}
			if intervals[i].Start == intervals[j].Start && intervals[i].End <= intervals[j].End {
				intervals[i].End = intervals[j].End
			}
			if intervals[i].Start < intervals[j].Start && intervals[i].End >= intervals[j].End {
				intervals[j] = intervals[i]
			}
			if intervals[i].Start < intervals[j].Start && intervals[i].End >= intervals[j].Start && intervals[i].End < intervals[j].End {
				intervals[i].End = intervals[j].End
				intervals[j].Start = intervals[i].Start
			}
		}
		intervalM[intervals[i]] = true
	}
	// 3）取值
	for k, _ := range intervalM {
		res = append(res, k)
	}
	return res
}

func main() {
	// [[1,3],[2,6],[8,10],[15,18]]
	// [1,4],[4,5]
	//i := []Interval{
	//	Interval{1,3},
	//	Interval{2,6},
	//	Interval{8,10},
	//	Interval{15,18},
	//}[[1,4],[0,5]]
	i := []Interval{
		Interval{1,4},
		Interval{1,3},
	}
	fmt.Println(merge(i))
}