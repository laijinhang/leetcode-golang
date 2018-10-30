package main

import (
	"sort"
)

type ds []string

func (d ds)Len() int {
	return len(d)
}

func (d ds)Less(i, j int) bool {    // 返回false表示交换
	if (len(d[i]) < len(d[j])) || (len(d[i]) == len(d[j]) && d[i] > d[j]) {
		return false
	}
	return true
}

func (d ds)Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func findLongestWord(s string, d []string) string {
	// 1）对字典进行排序，排序规则是长度长的在前面，长度相同则字典序小的在前面
	sort.Sort(ds(d))
	// 2）寻找符合的字符串
	res := ""
	for i := 0;i < len(d);i++ {
		j, k := 0, 0
		for ;j < len(s) && k < len(d[i]);j++ {
			if d[i][k] == s[j] {
				k++
			}
		}
		if k == len(d[i]) {
			res = d[i]
			break
		}
	}
	return res
}