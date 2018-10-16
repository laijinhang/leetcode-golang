package main

import (
	"sort"
)

type newPeople [][]int

func reconstructQueue(people [][]int) [][]int {
	// 重排
	sort.Sort(newPeople(people))
	res := [][]int{}
	for i := 0; i < len(people); i++ {
		if people[i][1] <= len(people) {
			res = append(res[:people[i][1]], append([][]int{people[i]}, res[people[i][1]:]...)...)
			continue
		}
		res = append(res, people[i])
	}
	return res
}

// 从大到小，相同则排在前面人数少的在前面
func (p newPeople) Less(i, j int) bool {
	if p[i][0] > p[j][0] {
		return true
	}
	if p[i][0] == p[j][0] && p[i][1] < p[j][1] {
		return true
	}
	return false
}

func (p newPeople) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p newPeople) Len() int {
	return len(p)
}
