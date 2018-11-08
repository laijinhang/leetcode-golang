package main

import "strings"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	A += A
	if strings.Index(A, B) == -1 {
		return false
	}
	return true
}