package main

func reverseString(s string) string {
	temp := []byte(s)
	for i, j := 0, len(temp) - 1;i < j;i, j = i + 1, j - 1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	s = string(temp)
	return s
}
