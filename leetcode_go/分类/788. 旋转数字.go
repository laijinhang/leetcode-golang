package main

func rotatedDigits(N int) int {
	num := 0
	n := 0
	for i := 2;i <= N;i++ {
		t := i
		n = 0
		for ;t != 0;t /= 10 {
			if t % 10 == 3 || t % 10 == 4 || t % 10 == 7 {
				break
			} else if t % 10 == 2 || t % 10 == 5 || t % 10 == 6 || t % 10 == 9 {
				n++
			}
		}
		if t == 0 && n != 0 {
			num++
		}
	}
	return num
}