package main

import (
	"sort"
)

type ByteSlice []byte

func (s ByteSlice) Len() int {
	return len(s)
}
func (s ByteSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByteSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func isAnagram(s string, t string) bool {
	temps := ByteSlice(s)
	tempt := ByteSlice(t)
	sort.Sort(temps)
	sort.Sort(tempt)
	if string(temps) != string(tempt) {
		return false
	}
	return true
}

func main()  {
	isAnagram("dsadas", "dxicaoudsoi")
}