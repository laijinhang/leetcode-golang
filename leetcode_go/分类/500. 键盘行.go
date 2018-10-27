package main

var f = map[byte]bool {
	'Q':true,'q':true,
	'W':true,'w':true,
	'E':true,'e':true,
	'R':true,'r':true,
	'T':true,'t':true,
	'Y':true,'y':true,
	'U':true,'u':true,
	'I':true,'i':true,
	'O':true,'o':true,
	'P':true,'p':true,
}
var t1 = map[byte]bool {
	'A':true,'a':true,
	'S':true,'s':true,
	'D':true,'d':true,
	'F':true,'f':true,
	'G':true,'g':true,
	'H':true,'h':true,
	'J':true,'j':true,
	'K':true,'k':true,
	'L':true,'l':true,
}
var t2 = map[byte]bool{
	'Z':true,'z':true,
	'X':true,'x':true,
	'C':true,'c':true,
	'V':true,'v':true,
	'B':true,'b':true,
	'N':true,'n':true,
	'M':true,'m':true,
}

func findWords(words []string) []string {
	res := []string{}
	for i, j := 0, 0;i < len(words);i++ {
		if f[words[i][0]] {
			for j = 0;j < len(words[i]);i++ {
				if f[words[i][j]] == false {
					break
				}
			}
			if j == len(words[i]) {
				res = append(res, words[i])
			}
		} else if t1[words[i][0]] {
			for j = 0;j < len(words[i]);i++ {
				if t1[words[i][j]] == false {
					break
				}
			}
			if j == len(words[i]) {
				res = append(res, words[i])
			}
		} else if t2[words[i][0]] {
			for j = 0;j < len(words[i]);i++ {
				if t2[words[i][j]] == false {
					break
				}
			}
			if j == len(words[i]) {
				res = append(res, words[i])
			}
		}
	}
	return res
}