package main

import (
	"strings"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for i := 0;i < len(cpdomains);i++ {
		tempv := strings.Split(cpdomains[i], " ")
		temps := strings.Split(tempv[1], ".")
		v, _ := strconv.Atoi(tempv[0])
		if len(temps) == 2 {
			m[temps[1]] += v
			m[tempv[1]] += v
		} else {
			m[temps[2]] += v
			m[temps[1]+"."+temps[2]] += v
			m[tempv[1]] += v
		}
	}
	res := []string{}
	for k, v := range m {
		res = append(res, strconv.Itoa(v) + " " + k)
	}
	return res
}