package main

func findTheDifference(s string, t string) byte {
	s1 := make([]int, 27)
	t1 := make([]int, 27)
	for i := 0;i < len(s);i++ {
		s1[s[i]-'a']++
	}
	for i := 0;i < len(t);i++ {
		t1[t[i]-'a']++
	}
	res := byte(0)
	for i := 0;i < 27;i++ {
		if s1[i] != t1[i] {
			res = byte(i) + 'a'
		}
	}
	return byte(res)
}