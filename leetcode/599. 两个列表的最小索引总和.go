package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	ml := make(map[string]int)
	for i := 0;i < len(list1);i++ {
		ml[list1[i]] = i
	}
	res := []string{}
	minIndex := math.MaxInt32
	for i := 0;i < len(list2);i++ {
		if _, ok := ml[list2[i]];ok && ml[list2[i]] + i <= minIndex {
			if minIndex == ml[list2[i]] + i {
				res = append(res, list2[i])
			} else {
				minIndex = ml[list2[i]] + i
				res = []string{list2[i]}
			}
		}
	}
	return res
}