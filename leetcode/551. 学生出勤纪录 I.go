package main

import "strings"

func checkRecord(s string) bool {
	if !(strings.Count(s, "A") <= 1 && strings.Count(s, "LLL") == 0) {
		return false
	}
	return true
}