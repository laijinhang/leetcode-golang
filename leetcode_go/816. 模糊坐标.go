package main

import "fmt"

func ambiguousCoordinates(S string) []string {
	res := []string{}
	strs := make(map[int]string)
	s := S[1:len(S)-1]
	for i := 1;i < len(s);i++ {
		s1 := s[:i]
		s2 := s[i:]
		for i1 := 0;k1 < i;k1++ {
			for i2 := i;k2 < len(s);k2++ {

				if i1 > len(s1) || i2 > len(s2) || (len(s1) > 1 && s1[0] == '0') || (len(s2) > 1 && s2[0] == '0') {
					return
				}
				if !((len(s1[:i1]) > 1 && s1[0] == '0') || zreoStr(s1[i1:]) ||
					(len(s2[:i2]) > 1 && s2[0] == '0') || 	zreoStr(s2[i2:])) {
					str := ""
					if len(s1) > i1 {
						str += s1[:i1] + "." + s1[i1:] + ", "
					} else {
						str += s1 + ", "
					}
					if len(s2) > i2 {
						str += s2[:i2] + "." + s2[i2:]
					} else {
						str += s2
					}
					strM[len(strM)] = "(" + str + ")"
				}
			}
		}
	}
	for _, v := range strs {
		res = append(res, v)
	}
	return res
}

func composition(s1, s2 string, i1, i2 int, strM map[int]string) {
	if i1 > len(s1) || i2 > len(s2) || (len(s1) > 1 && s1[0] == '0') || (len(s2) > 1 && s2[0] == '0') {
		return
	}
	if !((len(s1[:i1]) > 1 && s1[0] == '0') || zreoStr(s1[i1:]) ||
		(len(s2[:i2]) > 1 && s2[0] == '0') || 	zreoStr(s2[i2:])) {
		str := ""
		if len(s1) > i1 {
			str += s1[:i1] + "." + s1[i1:] + ", "
		} else {
			str += s1 + ", "
		}
		if len(s2) > i2 {
			str += s2[:i2] + "." + s2[i2:]
		} else {
			str += s2
		}
		strM[len(strM)] = "(" + str + ")"
	}
	composition(s1, s2, i1 + 1, i2, strM)
	composition(s1, s2, i1, i2 + 1, strM)
	composition(s1, s2, i1 + 1, i2 + 1, strM)
}

func zreoStr(s string) bool {
	if len(s) > 0 && s[len(s)-1] == '0' {
		return true
	}
	return false
}

func main()  {
	fmt.Println(ambiguousCoordinates("(3599191175)"))
}