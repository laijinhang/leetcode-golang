package main

func judgeCircle(moves string) bool {
	a := [2]int{}
	mLen := len(moves)
	for i := 0;i < mLen;i++ {
		switch moves[i] {
		case 'R':a[0]++
		case 'L':a[0]--
		case 'U':a[1]++
		case 'D':a[1]--
		}
	}
	if a[0] != 0 || a[1] != 0 {
		return false
	}
	return true
}